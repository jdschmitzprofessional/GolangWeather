package Structs

import "time"

type WeatherStation struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Properties struct {
		Data    string    `json:"data"`
		ID      string    `json:"id"`
		Site    string    `json:"site"`
		Prior   int       `json:"prior"`
		ObsTime time.Time `json:"obsTime"`
		Temp    float64   `json:"temp"`
		Dewp    float64   `json:"dewp"`
		Wspd    int       `json:"wspd"`
		Wdir    int       `json:"wdir"`
		Ceil    int       `json:"ceil"`
		Cover   string    `json:"cover"`
		CldCvg1 string    `json:"cldCvg1"`
		CldBas1 string    `json:"cldBas1"`
		Visib   float64   `json:"visib"`
		Fltcat  string    `json:"fltcat"`
		Altim   float64   `json:"altim"`
		Slp     float64   `json:"slp"`
		Wx      string    `json:"wx"`
		RawOb   string    `json:"rawOb"`
	} `json:"properties"`
}
