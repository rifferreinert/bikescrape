package divvy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rifferreinert/bikescrape/pkg/divvy/models"
)

func GetStations() ([]models.Station, error) {
	res, err := http.Get("https://feeds.divvybikes.com/stations/stations.json")
	if err != nil {
		return nil, fmt.Errorf("Error retrieving divvy data: %v", err)
	}

	defer res.Body.Close()
	stationList := models.Stations{}
	jsonResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading divvy response: %v", err)
	}
	if err = json.Unmarshal(jsonResponse, &stationList); err != nil {
		return nil, fmt.Errorf("Error parsing divvy JSON response: %v", err)
	}

	return stationList, nil
}
