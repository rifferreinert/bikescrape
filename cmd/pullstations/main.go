package main

import (
	"fmt"

	"github.com/getsentry/raven-go"

	"github.com/rifferreinert/bikescrape"
	"github.com/rifferreinert/bikescrape/pkg/divvy"
	"github.com/rifferreinert/bikescrape/pkg/divvy/models"
	"github.com/rifferreinert/bikescrape/pkg/logs"
)

func main() {
	raven.CapturePanicAndWait(
		run,
		bikescrape.RavenContext,
	)
}

func run() {
	config := models.DBConfig{
		Username: bikescrape.Username,
		Password: bikescrape.Password,
		Host:     bikescrape.Host,
		DBName:   bikescrape.DBName,
	}
	db, err := models.NewDB(&config)

	if err != nil {
		errMsg := fmt.Errorf("error opening DB: %v", err)
		logs.Fatal(errMsg)
	}
	defer db.Close()

	stations, err := divvy.GetStations()
	if err != nil {
		msg := fmt.Errorf("error querying DIVVY API: %v", err)
		logs.Fatal(msg)
	}

	tx := db.Begin()
	for _, station := range stations {
		if err := tx.Create(&station).Error; err != nil {
			tx.Rollback()
			msg := fmt.Errorf("Error inserting station record: %v", err)
			logs.Fatal(msg)
		}
	}
	tx.Commit()
}
