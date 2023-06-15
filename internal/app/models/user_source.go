package models

import "time"

type UserSource struct {
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserID   int64     `gorm:"column:user_id;NOT NULL"`
	SourceID int64     `gorm:"column:source_id;NOT NULL"`
	Ctime    time.Time `gorm:"column:ctime;default:current_timestamp();NOT NULL"`
	Utime    time.Time `gorm:"column:utime;default:current_timestamp();NOT NULL"`
}
