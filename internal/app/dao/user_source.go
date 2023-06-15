package dao

import (
	"gorm.io/gorm"
	"iptv-checker/internal/app/dto"
	"iptv-checker/internal/app/models"
	"iptv-checker/internal/app/models/database"
	"iptv-checker/pkg/log"
)

type UserSource interface {
	Add() (int32, error)
	Delete() (int32, error)
	Update() (int32, error)
	Select() []*models.UserSource
}

type UserSourceDAO struct {
	db *gorm.DB
}

func (dao *UserSourceDAO) Add() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *UserSourceDAO) Delete() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *UserSourceDAO) Update() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *UserSourceDAO) Select(cond map[string]interface{}) []*models.UserSource {
	return nil
}

func (dao *UserSourceDAO) ListUserSource(cond map[string]interface{}) ([]*dto.UserSourceDTO, error) {

	usDTO := make([]*dto.UserSourceDTO, 0)

	rows, err := dao.db.Table("user").
		Select("user.user_id, user.nick_name, " +
			"live_source.name as source_name, " +
			"live_source.source as source_url, " +
			"live_source.data as source_data, " +
			"live_source.ctime, " +
			"live_source.utime").
		Where(cond).
		Joins("left join user_source on user_source.user_id = user.user_id").
		Joins("left join live_source on live_source.id = user_source.source_id").
		Rows()
	if err != nil {
		log.Logger.Error(err)
		return nil, err
	}

	for rows.Next() {
		t := dto.UserSourceDTO{}
		err := dao.db.ScanRows(rows, &t)
		if err != nil {
			continue
		}
		usDTO = append(usDTO, &t)
	}

	return usDTO, nil

}

func NewUserSourceDAO() *UserSourceDAO {
	return &UserSourceDAO{
		db: database.DB,
	}
}
