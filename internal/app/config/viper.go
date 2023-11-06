package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitLoadConfiguration() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	viper.AddConfigPath(dir + "/cmd")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}

func InitLoadConfigurationMain() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	viper.AddConfigPath(dir)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
