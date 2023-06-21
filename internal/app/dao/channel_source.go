package dao

import (
	"gorm.io/gorm"
	"iptv-checker/core/log"
	"iptv-checker/internal/app/models"
	"iptv-checker/internal/app/models/database"
)

type ChannelSource interface {
	Add(models.ChannelSource) (int64, error, int64)
	Delete() (int32, error)
	Update() (int32, error)
	Select() []*models.ChannelSource
	Find(map[string]interface{}) []*models.ChannelSource
}

type ChannelSourceDAO struct {
	db *gorm.DB
}

var tblChannelSource = "channel_source"

func (dao *ChannelSourceDAO) Add(cs models.ChannelSource) (int64, error, int64) {
	//TODO implement me
	panic("implement me")
}

func (dao *ChannelSourceDAO) Delete() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *ChannelSourceDAO) Update() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *ChannelSourceDAO) Select() []*models.ChannelSource {
	//TODO implement me
	panic("implement me")
}

func (dao *ChannelSourceDAO) Find(cond map[string]interface{}) []*models.ChannelSource {
	var cs []*models.ChannelSource
	if err := dao.db.Table(tblChannelSource).Where(cond).Find(&cs).Error; err != nil {
		log.Logger.Error(err)
	}
	return cs
}

func NewChannelSourceDAO() *ChannelSourceDAO {
	return &ChannelSourceDAO{
		db: database.DB,
	}
}
