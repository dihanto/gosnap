package config

import (
	"log"

	"github.com/spf13/viper"
)

func ViperReadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("cmd")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
}
