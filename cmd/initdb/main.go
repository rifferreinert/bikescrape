package main

import (
	"fmt"
	"log"

	"github.com/getsentry/raven-go"

	"github.com/rifferreinert/bikescrape"
	"github.com/rifferreinert/bikescrape/pkg/divvy/models"
)

func main() {
	db, err := models.NewDB(&bikescrape.DbConfig)
	if err != nil {
		msg := fmt.Errorf("Error opening db: %v", err)
		raven.CaptureErrorAndWait(msg, bikescrape.RavenContext)
		log.Fatal(msg)
	}
	if err := db.AutoMigrate(&models.Station{}).Error; err != nil {
		msg := fmt.Errorf("error migrating db: %v", err)
		raven.CaptureErrorAndWait(msg, bikescrape.RavenContext)
		log.Fatal(msg)
	}
	defer db.Close()
}
