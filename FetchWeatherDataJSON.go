package main

import (
	"AviaWeather/Structs"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/jsonq"
	"log"
	"net/http"
)

func main() {
	var err error
	data := map[string]interface{}{}

	//targetStations := [5]string{"KPHX", "KSDL", "KLUF", "KGEU", "KDVT"}
	resp, err := http.Get("https://aviationweather.gov/cgi-bin/json/MetarJSON.php")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	jq := jsonq.NewQuery(data)
	station_array, err := jq.Array("features")
	if err != nil {
		log.Fatal(err)
	}
	arr := make([]Structs.WeatherStation, len(station_array))
	fmt.Println(len(arr))
	for stat := range station_array {
		if stat == 0 {
			continue
		}
		json_string, _ := json.Marshal(station_array[stat])
		current_struct := Structs.WeatherStation{}
		err = json.Unmarshal(json_string, &current_struct)
		if err != nil {
			log.Fatal(err)
		}
		if current_struct.Properties.ID == "KPHX" {
			fmt.Println(current_struct.Properties)
		} else {
			fmt.Println(current_struct.Properties.ID)
		}

	}

}
