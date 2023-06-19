package dao

import (
	"gorm.io/gorm"
	"iptv-checker/core/log"
	"iptv-checker/internal/app/models"
	"iptv-checker/internal/app/models/database"
)

type LiveSource interface {
	Add(models.LiveSource) (int64, error, int64)
	Delete() (int32, error)
	Update() (int32, error)
	Select() ([]models.LiveSource, error)
}

type LiveSourceDAO struct {
	db *gorm.DB
}

var tblLiveSource = "live_source"

func (dao *LiveSourceDAO) Add(ls models.LiveSource) (int64, error, int64) {
	result := dao.db.Table(tblLiveSource).Create(&ls)
	return result.RowsAffected, result.Error, ls.ID
}

func (dao *LiveSourceDAO) AddWithTX(tx *gorm.DB, ls models.LiveSource) (*gorm.DB, int64, error, int64) {
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
	result := tx.Table(tblLiveSource).Create(&ls)
	if result.Error != nil {
		tx.Rollback()
		return tx, 0, result.Error, 0
	}

	return tx, result.RowsAffected, nil, ls.ID
}

func (dao *LiveSourceDAO) Delete() (int32, error) {
	//TODO implement me
	panic("implement me")
}
func (dao *LiveSourceDAO) DeleteWithTX(tx *gorm.DB, liveSourceID []string) (*gorm.DB, int64, error) {
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

	result := tx.Table(tblLiveSource).Where("id in ?", liveSourceID).Delete(&models.LiveSource{})
	if result.Error != nil {
		tx.Rollback()
		return tx, 0, result.Error
	}
	return tx, result.RowsAffected, nil
}

func (dao *LiveSourceDAO) Update() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *LiveSourceDAO) Select() ([]models.LiveSource, error) {
	//TODO implement me
	panic("implement me")
}

func NewLiveSourceDAO() *LiveSourceDAO {
	return &LiveSourceDAO{
		db: database.DB,
	}
}
