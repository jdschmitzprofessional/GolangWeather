package models

import (
	. "fmt"
	. "strings"
	. "time"
)

// WeatherStation is a container for weather data
type WeatherStation struct {
	Type  string     `json:"type"`
	ID    string     `json:"id"`
	Props Properties `json:"properties"`
}

// ToString prints the string representation of a WeatherStation struct
func (w *WeatherStation) ToString() (str string) {
	builder := Builder{}

	builder.WriteString("WeatherStation[ ")
	builder.WriteString(Sprintf("ID: %s, ", w.ID))
	builder.WriteString(Sprintf("Type: %s, ", w.Type))
	builder.WriteString(Sprintf("Props: %s", w.Props.ToString()))
	builder.WriteString(" ]")

	return builder.String()
}

// Properties is a container for WeatherStation props
type Properties struct {
	Data    string  `json:"data"`
	ID      string  `json:"id"`
	Site    string  `json:"site"`
	Prior   int     `json:"prior"`
	ObsTime Time    `json:"obsTime"`
	Temp    float64 `json:"temp"`
	Dewp    float64 `json:"dewp"`
	Wspd    int     `json:"wspd"`
	Wdir    int     `json:"wdir"`
	Ceil    int     `json:"ceil"`
	Cover   string  `json:"cover"`
	CldCvg1 string  `json:"cldCvg1"`
	CldBas1 string  `json:"cldBas1"`
	Visib   float64 `json:"visib"`
	Fltcat  string  `json:"fltcat"`
	Altim   float64 `json:"altim"`
	Slp     float64 `json:"slp"`
	Wx      string  `json:"wx"`
	RawOb   string  `json:"rawOb"`
}

// ToString prints the string representation of a Properties struct
func (p *Properties) ToString() (str string) {
	builder := Builder{}

	builder.WriteString("Properties[ ")
	builder.WriteString(Sprintf("Data: %s, ", p.Data))
	builder.WriteString(Sprintf("ID: %s, ", p.ID))
	builder.WriteString(Sprintf("Site: %s, ", p.Site))
	builder.WriteString(Sprintf("Prior: %d, ", p.Prior))
	builder.WriteString(Sprintf("ObsTime: %s, ", p.ObsTime.String()))
	builder.WriteString(Sprintf("Temp: %.2f, ", p.Temp))
	builder.WriteString(Sprintf("Dewp: %.2f, ", p.Dewp))
	builder.WriteString(Sprintf("Wspd: %d, ", p.Wspd))
	builder.WriteString(Sprintf("Wdir: %d, ", p.Wdir))
	builder.WriteString(Sprintf("Ceil: %d, ", p.Ceil))
	builder.WriteString(Sprintf("Cover: %s, ", p.Cover))
	builder.WriteString(Sprintf("CldCvg1: %s, ", p.CldCvg1))
	builder.WriteString(Sprintf("CldBas1: %s, ", p.CldBas1))
	builder.WriteString(Sprintf("Visib: %.2f, ", p.Visib))
	builder.WriteString(Sprintf("Fltcat: %s, ", p.Fltcat))
	builder.WriteString(Sprintf("Altim: %.2f, ", p.Altim))
	builder.WriteString(Sprintf("Slp: %.2f, ", p.Slp))
	builder.WriteString(Sprintf("Wx: %s, ", p.Wx))
	builder.WriteString(Sprintf("RawOb: %s", p.RawOb))
	builder.WriteString(" ]")

	return builder.String()
}
