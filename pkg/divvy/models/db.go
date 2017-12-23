package models

import (
	"fmt"

	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rifferreinert/bikescrape"
)

type DBConfig struct {
	Username string
	Password string
	DBName   string
	Host     string
}

func NewDB(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s user=%s dbname=%s password=%s", config.Host, config.Username, config.DBName, config.Password))
	if err != nil {
		err = fmt.Errorf("Error connecting to DB: %v", err)
		raven.CaptureErrorAndWait(err, bikescrape.RavenContext)
		return nil, err
	}

	return db, nil
}
