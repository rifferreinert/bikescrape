package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Station struct {
	gorm.Model
	DivvyID               *int
	StationName           *string
	AvailableDocks        *int
	TotalDocks            *int
	Latitude              *float64
	Longitude             *float64
	StatusValue           *string
	StatusKey             *int
	AvailableBikes        *int
	Status                *string
	StAddress1            *string
	StAddress2            *string
	City                  *string
	PostalCode            *string
	Location              *string
	Altitude              *string
	TestStation           *bool
	LastCommunicationTime *time.Time
	LandMark              *string
	IsRenting             *bool
}

type Stations []Station

func (s *Stations) UnmarshalJSON(data []byte) error {
	response := Response{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return fmt.Errorf("error unmarshalling stations: %v", err)
	}
	for _, stationResponse := range response.StationBeanList {
		station := Station{}

		station.DivvyID = stationResponse.ID
		station.StationName = stationResponse.StationName
		station.AvailableDocks = stationResponse.AvailableDocks
		station.TotalDocks = stationResponse.TotalDocks
		station.Latitude = stationResponse.Latitude
		station.Longitude = stationResponse.Longitude
		station.StatusValue = stationResponse.StatusValue
		station.StatusKey = stationResponse.StatusKey
		station.AvailableBikes = stationResponse.AvailableBikes
		station.Status = stationResponse.Status
		station.StAddress1 = stationResponse.StAddress1
		station.StAddress2 = stationResponse.StAddress2
		station.City = stationResponse.City
		station.PostalCode = stationResponse.PostalCode
		station.Location = stationResponse.Location
		station.Altitude = stationResponse.Altitude
		station.TestStation = stationResponse.TestStation
		if stationResponse.LastCommunicationTime != nil {
			t, err := time.Parse("2006-01-02 15:04:05", *stationResponse.LastCommunicationTime)
			if err != nil {
				return fmt.Errorf("error parsing time from DIVVY API: %v", err)
			}
			station.LastCommunicationTime = &t
		}
		station.LandMark = stationResponse.LandMark
		station.IsRenting = stationResponse.IsRenting

		*s = append(*s, station)
	}

	return nil
}
