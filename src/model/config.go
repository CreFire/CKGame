package model

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	Port string
}

var Config config

func LoadConfig() {
	viper.SetConfigName("default")
	// 添加配置文件所在的路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	// 尝试读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(Config)
	if err != nil {
		log.Fatal("Error Unmarshal config", err)
		return
	}
}
