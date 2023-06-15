package models

import "time"

type User struct {
	ID       int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserID   int64     `gorm:"column:user_id;NOT NULL"`
	NickName string    `gorm:"column:nick_name;default:;NOT NULL"`
	SourceID int64     `gorm:"column:source_id;default:0;NOT NULL"`
	Ctime    time.Time `gorm:"column:ctime;default:current_timestamp();NOT NULL"`
	Utime    time.Time `gorm:"column:utime;default:current_timestamp();NOT NULL"`
}
