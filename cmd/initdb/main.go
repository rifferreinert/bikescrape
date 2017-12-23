package main

import (
	"fmt"

	"github.com/rifferreinert/bikescrape"
	"github.com/rifferreinert/bikescrape/pkg/divvy/models"
	"github.com/rifferreinert/bikescrape/pkg/logs"
)

func main() {
	config := models.DBConfig{
		Username: bikescrape.Username,
		Password: bikescrape.Password,
		Host:     bikescrape.Host,
		DBName:   bikescrape.DBName,
	}
	db, err := models.NewDB(&config)
	if err != nil {
		msg := fmt.Errorf("Error opening db: %v", err)
		logs.Fatal(msg)
	}
	if err := db.AutoMigrate(&models.Station{}).Error; err != nil {
		msg := fmt.Errorf("error migrating db: %v", err)
		logs.Fatal(msg)
	}
	defer db.Close()
}
