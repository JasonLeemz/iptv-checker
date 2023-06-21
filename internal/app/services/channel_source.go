package services

import (
	"iptv-checker/internal/app/dao"
	"iptv-checker/internal/app/models"
)

type ChannelSourceService struct {
	ChannelSourceDAO *dao.ChannelSourceDAO
}

var theChannelSourceService = new(ChannelSourceService)

func NewChannelSourceService() *ChannelSourceService {
	// todo use a more elegant approach to implement
	if theChannelSourceService.ChannelSourceDAO == nil {
		theChannelSourceService.ChannelSourceDAO = dao.NewChannelSourceDAO()
	}
	return theChannelSourceService
}

func (s *ChannelSourceService) QueryChannelSource(sourceID []int64) map[int64][]*models.ChannelSource {
	cond := map[string]interface{}{
		"source_id": sourceID,
	}
	channels := s.ChannelSourceDAO.Find(cond)

	cs := make(map[int64][]*models.ChannelSource)
	for _, channel := range channels {
		if cs[channel.SourceId] == nil {
			cs[channel.SourceId] = make([]*models.ChannelSource, 0)
		}
		cs[channel.SourceId] = append(cs[channel.SourceId], channel)
	}

	return cs
}
