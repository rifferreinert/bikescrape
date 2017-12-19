package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBConfig struct {
	Username string
	Password string
	DBName   string
	Host     string
}

func NewDB(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", config.Host, config.Username, config.DBName, config.Password))
	if err != nil {
		err = fmt.Errorf("Error connecting to DB: %v", err)
		return nil, err
	}

	return db, nil
}
