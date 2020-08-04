package util

import (
	"fmt"
	"github.com/cfi2017/sheep-bot/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	db *gorm.DB
)

func InitialisePersistence() (*gorm.DB, error) {
	var (
		host     = viper.GetString("db.host")
		port     = viper.GetInt("db.port")
		database = viper.GetString("db.database")
		username = viper.GetString("db.username")
		password = viper.GetString("db.password")
	)

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.UserGuildRole{})

	return db, nil
}

func SetDatabase(database *gorm.DB) {
	db = database
}

func GetDatabase() *gorm.DB {
	return db
}
