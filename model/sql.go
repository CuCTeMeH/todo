package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var client *gorm.DB

func initDB() error {
	if client != nil {
		return nil
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	db := viper.GetString("DB_NAME")
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")

	debug := viper.GetBool("DEBUG")

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, db)
	connection, err := gorm.Open("mysql", dbSource)

	if err != nil {
		panic("failed to connect database")
	}

	if debug == true {
		connection.LogMode(true)
	}
	client = connection
	return connection.DB().Ping()
}

func Client() *gorm.DB {
	if client == nil {
		err := initDB()
		if err != nil {
			panic(err)
		}
	}

	return client
}
