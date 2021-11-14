package main

import (
	"encoding/json"
	. "github.com/jdschmitzprofessional/GolangWeather/models"
	"github.com/jmoiron/jsonq"
	"io"
	"log"
	"net/http"
)

const ApiUrl = "https://aviationweather.gov/cgi-bin/json/MetarJSON.php"

func main() {
	var err error
	var data map[string]interface{}

	resp, err := http.Get(ApiUrl)
	LogError(err)

	err = DecodeResponse(resp, &data)
	LogError(err)

	jq := jsonq.NewQuery(data)
	stationArray, err := jq.Array("features")
	LogError(err)

	arr := make([]*WeatherStation, len(stationArray))

	for index := range stationArray {
		newStation, err := JsonToStation(stationArray[index])
		LogError(err)

		arr[index] = newStation

		println(newStation.ToString())
	}
}

// DecodeResponse Decodes a response body to a map
// returns error if err occurs during decode or closing body
func DecodeResponse(resp *http.Response, data *map[string]interface{}) (err error) {
	defer func(Body io.ReadCloser) { LogError(Body.Close()) }(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&data)
	return err
}

// JsonToStation converts json data to Station
// returns error if err occurs conversion
func JsonToStation(v interface{}) (station *WeatherStation, err error) {
	if jsonString, err := json.Marshal(v); err != nil {
		return nil, err
	} else if err = json.Unmarshal(jsonString, &station); err != nil {
		return nil, err
	}
	return station, err
}

// LogError logs any incoming errors
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
