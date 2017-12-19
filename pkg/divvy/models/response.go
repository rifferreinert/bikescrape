package models

type StationResponse struct {
	ID                    *int     `json:"id"`
	StationName           *string  `json:"stationName"`
	AvailableDocks        *int     `json:"availableDocks"`
	TotalDocks            *int     `json:"totalDocks"`
	Latitude              *float64 `json:"latitude"`
	Longitude             *float64 `json:"longitude"`
	StatusValue           *string  `json:"statusValue"`
	StatusKey             *int     `json:"statusKey"`
	AvailableBikes        *int     `json:"availableBikes"`
	Status                *string  `json:"status"`
	StAddress1            *string  `json:"stAddress1"`
	StAddress2            *string  `json:"stAddress2"`
	City                  *string  `json:"city"`
	PostalCode            *string  `json:"postalCode"`
	Location              *string  `json:"location"`
	Altitude              *string  `json:"altitude"`
	TestStation           *bool    `json:"testStation"`
	LastCommunicationTime *string  `json:"lastCommunicationTime"`
	LandMark              *string  `json:"landMark"`
	IsRenting             *bool    `json:"is_renting"`
}

type Response struct {
	ExecutionTime   string            `json:"executionTime"`
	StationBeanList []StationResponse `json:"stationBeanList"`
}
