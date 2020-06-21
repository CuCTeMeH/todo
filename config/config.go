package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Settings *viper.Viper

func InitConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	Settings = viper.GetViper()
}
