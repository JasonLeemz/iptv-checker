package services

import (
	"iptv-checker/internal/app/dao"
	"iptv-checker/internal/app/dto"
)

type UserSourceService struct {
	UserSourceDAO *dao.UserSourceDAO
}

func NewUserSourceService(userSourceDAO *dao.UserSourceDAO) *UserSourceService {
	return &UserSourceService{
		UserSourceDAO: userSourceDAO,
	}
}

func (s *UserSourceService) GetUserSource(userID int64) ([]*dto.UserSourceDTO, error) {
	cond := map[string]interface{}{
		"user.user_id": userID,
	}
	return s.UserSourceDAO.ListUserSource(cond)
}
