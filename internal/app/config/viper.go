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
		viper.AddConfigPath(dir)
		err2 := viper.ReadInConfig()
		if err2 != nil {
			log.Fatalln(err2)
		}
	}
}
