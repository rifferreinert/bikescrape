package main

import (
	"fmt"
	"log"

	"github.com/getsentry/raven-go"

	"github.com/rifferreinert/bikescrape"
	"github.com/rifferreinert/bikescrape/pkg/divvy"
	"github.com/rifferreinert/bikescrape/pkg/divvy/models"
)

func main() {
	raven.CapturePanicAndWait(
		run,
		bikescrape.RavenContext,
	)
}

func run() {
	db, err := models.NewDB(&bikescrape.DbConfig)
	if err != nil {
		errMsg := fmt.Errorf("error opening DB: %v", err)
		raven.CaptureErrorAndWait(errMsg, bikescrape.RavenContext)
		log.Fatal(errMsg)
	}
	defer db.Close()

	stations, err := divvy.GetStations()
	if err != nil {
		msg := fmt.Errorf("error querying DIVVY API: %v", err)
		raven.CaptureErrorAndWait(msg, bikescrape.RavenContext)
		log.Fatal(msg)
	}

	tx := db.Begin()
	for _, station := range stations {
		if err := tx.Create(&station).Error; err != nil {
			tx.Rollback()
			msg := fmt.Errorf("Error inserting station record: %v", err)
			raven.CaptureErrorAndWait(msg, bikescrape.RavenContext)
			log.Fatal(msg)
		}
	}
	tx.Commit()
}
