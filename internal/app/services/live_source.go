package services

import (
	"iptv-checker/internal/app/dao"
	"iptv-checker/internal/app/dto"
)

type LiveSourceService struct {
	LiveSourceDAO *dao.LiveSourceDAO
}

func NewLiveSourceService(liveSourceDAO *dao.LiveSourceDAO) *LiveSourceService {
	return &LiveSourceService{
		LiveSourceDAO: liveSourceDAO,
	}
}

func (s *LiveSourceService) GetLiveSourceByUserID(userID int64) (*dto.UserSourceDTO, error) {

	return &dto.UserSourceDTO{}, nil
}
