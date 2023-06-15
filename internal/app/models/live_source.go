package models

import "time"

type LiveSource struct {
	ID     int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name   string    `gorm:"column:name;NOT NULL"`
	Source string    `gorm:"column:source;NOT NULL"`
	Data   string    `gorm:"column:data;default:NULL"`
	Ctime  time.Time `gorm:"column:ctime;default:current_timestamp();NOT NULL"`
	Utime  time.Time `gorm:"column:utime;default:current_timestamp();NOT NULL"`
}
