package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 设置配置文件名称和路径
	viper.SetConfigName("app.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read config file:" + err.Error())
	}

	// 读取配置项
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")

	// 打印配置项值
	fmt.Println("Host:", host)
	fmt.Println("Port:", port)
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
}
