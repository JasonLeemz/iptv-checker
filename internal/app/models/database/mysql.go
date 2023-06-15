package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iptv-checker/pkg/log"
)

var DB *gorm.DB

func DatabaseInit() {
	dsn := "root:root@tcp(mariadb.ybdx.xyz:3306)/iptv?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log.Glogger,
	})
	if err != nil {
		panic(err)
	}
}
