package dao

import (
	"gorm.io/gorm"
	"iptv-checker/internal/app/models"
	"iptv-checker/internal/app/models/database"
)

type LiveSource interface {
	Add() (int32, error)
	Delete() (int32, error)
	Update() (int32, error)
	Select() ([]models.LiveSource, error)
}

type LiveSourceDAO struct {
	db *gorm.DB
}

func (dao *LiveSourceDAO) Add() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *LiveSourceDAO) Delete() (int32, error) {
	//TODO implement me
	panic("implement me")
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
