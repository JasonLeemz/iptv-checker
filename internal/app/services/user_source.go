package services

import (
	"iptv-checker/core/log"
	"iptv-checker/internal/app/dao"
	"iptv-checker/internal/app/dto"
	"iptv-checker/internal/app/models"
	"time"
)

type UserSourceService struct {
	UserSourceDAO *dao.UserSourceDAO
	LiveSourceDAO *dao.LiveSourceDAO
}

var theUserSourceService = new(UserSourceService)

func NewUserSourceService() *UserSourceService {
	// todo use a more elegant approach to implement
	if theUserSourceService.UserSourceDAO == nil || theUserSourceService.LiveSourceDAO == nil {
		theUserSourceService.UserSourceDAO = dao.NewUserSourceDAO()
		theUserSourceService.LiveSourceDAO = dao.NewLiveSourceDAO()
	}
	return theUserSourceService
}

func (s *UserSourceService) GetUserSource(userID int64) ([]*dto.UserSourceDTO, error) {
	cond := map[string]interface{}{
		"user.user_id": userID,
	}
	us, err := s.UserSourceDAO.ListUserSource(cond)

	sourceID := make([]int64, 0)
	for _, source := range us {
		sourceID = append(sourceID, source.LiveSourceID)
	}

	ss := NewChannelSourceService()
	channelMap := ss.QueryChannelSource(sourceID)

	for _, source := range us {
		channels := channelMap[source.LiveSourceID]
		for _, channel := range channels {
			source.SourceData = append(source.SourceData, dto.ChannelSource{
				ChannelID:   channel.ID,
				ChannelName: channel.Name,
				ChannelUrl:  channel.Url,
				Ctime:       channel.Ctime.Format(time.RFC3339),
				Utime:       channel.Utime.Format(time.RFC3339),
			})
		}

	}
	return us, err
}

func (s *UserSourceService) AddSource(userID int64, sourceName, sourceUrl string) (int64, int64, error) {
	// 添加live_source
	ldata := models.LiveSource{
		Name:   sourceName,
		Source: sourceUrl,
	}
	tx, lAffected, err, sourceID := s.LiveSourceDAO.AddWithTX(nil, ldata)
	if err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		return lAffected, 0, err
	}

	// 添加user_source
	udata := models.UserSource{
		UserID:   userID,
		SourceID: sourceID,
	}

	tx, uAffected, err, _ := s.UserSourceDAO.AddWithTX(tx, udata)
	if err != nil {
		log.Logger.Error(err)
		return lAffected, uAffected, err
	}

	if tx != nil {
		err = tx.Commit().Error
	}
	return lAffected, uAffected, err
}

func (s *UserSourceService) DelSource(userSourceID, liveSourceID []string) error {
	// 删除user_source
	tx, _, err := s.UserSourceDAO.DeleteWithTX(nil, userSourceID)
	if err != nil {
		log.Logger.Error(err)
		return err
	}

	// 删除live_source
	_, _, err = s.LiveSourceDAO.DeleteWithTX(tx, liveSourceID)
	if err != nil {
		tx.Rollback()
		log.Logger.Error(err)
		return err
	}

	if tx != nil {
		err = tx.Commit().Error
	}
	return err
}
