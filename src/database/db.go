package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func InitDB(config Config) (*gorm.DB, error) {
	dsn := config.DBUsername + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
