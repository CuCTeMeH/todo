package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"todo/config"
)

var client *gorm.DB

func initDB() error {
	if client != nil {
		return nil
	}

	db := config.Settings.GetString("DB_NAME")
	username := config.Settings.GetString("DB_USERNAME")
	password := config.Settings.GetString("DB_PASSWORD")
	host := config.Settings.GetString("DB_HOST")
	port := config.Settings.GetInt("DB_PORT")

	debug := config.Settings.GetBool("DEBUG")

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
