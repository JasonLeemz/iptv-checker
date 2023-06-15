package utils

import (
	"os"
	"time"
)

func todayFilename() string {
	today := "logs/" + time.Now().Format("2006-01-02")
	return today + ".log"
}

func NewLogFile() *os.File {
	filename := todayFilename()
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
