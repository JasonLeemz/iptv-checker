package dao

import (
	"gorm.io/gorm"
	"iptv-checker/core/log"
	"iptv-checker/internal/app/dto"
	"iptv-checker/internal/app/models"
	"iptv-checker/internal/app/models/database"
)

type UserSource interface {
	Add(models.UserSource) (int64, error, int64)
	Delete() (int32, error)
	Update() (int32, error)
	Select() []*models.UserSource
	Find(map[string]interface{}) []*models.UserSource
}

type UserSourceDAO struct {
	db *gorm.DB
}

var tblUserSource = "user_source"

func (dao *UserSourceDAO) Add(us models.UserSource) (int64, error, int64) {
	result := dao.db.Table(tblUserSource).Create(&us)
	return result.RowsAffected, result.Error, us.ID
}

// AddWithTX return tx,add rows error id
func (dao *UserSourceDAO) AddWithTX(tx *gorm.DB, us models.UserSource) (*gorm.DB, int64, error, int64) {
	// 开始事务
	if tx == nil {
		tx = dao.db.Begin()
	}
	defer func() {
		if err := recover(); err != nil {
			log.Logger.Error(err)
			tx.Rollback()
		}
	}()

	result := tx.Table(tblUserSource).Create(&us)
	if result.Error != nil {
		tx.Rollback()
		return tx, 0, result.Error, 0
	}
	return tx, result.RowsAffected, nil, us.ID
}

func (dao *UserSourceDAO) Delete() (int32, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteWithTX return tx,delete rows error
func (dao *UserSourceDAO) DeleteWithTX(tx *gorm.DB, userSourceID []string) (*gorm.DB, int64, error) {
	// 开始事务
	if tx == nil {
		tx = dao.db.Begin()
	}
	defer func() {
		if err := recover(); err != nil {
			log.Logger.Error(err)
			tx.Rollback()
		}
	}()

	result := tx.Table(tblUserSource).Where("id in ?", userSourceID).Delete(&models.UserSource{})
	if result.Error != nil {
		tx.Rollback()
		return tx, 0, result.Error
	}
	return tx, result.RowsAffected, nil
}

func (dao *UserSourceDAO) Update() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *UserSourceDAO) Select() []*models.UserSource {
	return nil
}

func (dao *UserSourceDAO) Find(cond map[string]interface{}) []*models.UserSource {
	var us []*models.UserSource
	if err := dao.db.Table(tblUserSource).Where(cond).Find(&us).Error; err != nil {
		log.Logger.Error(err)
	}
	return us
}

func (dao *UserSourceDAO) ListUserSource(cond map[string]interface{}) ([]*dto.UserSourceDTO, error) {

	usDTO := make([]*dto.UserSourceDTO, 0)

	rows, err := dao.db.Table("user").
		Select("user.user_id, user.nick_name, " +
			"user_source.id as user_source_id, " +
			"live_source.id as live_source_id, " +
			"live_source.name as source_name, " +
			"live_source.source as source_url, " +
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
