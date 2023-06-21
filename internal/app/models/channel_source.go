package models

import "time"

type ChannelSource struct {
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	SourceId int64     `gorm:"column:source_id;NOT NULL"`
	Name     string    `gorm:"column:name;NOT NULL"`
	Url      string    `gorm:"column:url;NOT NULL"`
	Ctime    time.Time `gorm:"column:ctime;default:current_timestamp();NOT NULL"`
	Utime    time.Time `gorm:"column:utime;default:current_timestamp();NOT NULL"`
}
