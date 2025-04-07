package weather

import (
	"encoding/json"
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Yellow = "\033[33m"
)

func format_get(geoJson []byte) ([]CityInfo, error) {
	var geo []CityInfo
	err := json.Unmarshal([]byte(geoJson), &geo)
	if err != nil {
		return nil, err
	}
	return geo, nil

}

func ShowInfo(data []byte) {
	var result DataDisplay
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(Yellow + "══════════════════════════════" + Reset)
	fmt.Printf(Green+"🌍	Город: %s\n"+Reset, result.Name)
	fmt.Printf(Green+"	Широта: %f \n 	Долгота: %f \n"+Reset, result.Coord.Lat, result.Coord.Lon)
	fmt.Println(Blue + "══════════════════════════════" + Reset)
	fmt.Printf(Red+"🌦	Погода: %s\n"+Reset, result.Weather[0].Main)
	fmt.Println(Cyan + "══════════════════════════════" + Reset)
	fmt.Printf(Yellow+"🌡	Температура: %.2f°C\n"+Reset, result.Main.Temp)
	fmt.Println(Yellow + "══════════════════════════════" + Reset)
}

type CityInfo struct {
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type DataDisplay struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`

	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Base string `json:"base"`

	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`

	Visibility int `json:"visibility"`

	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`

	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain,omitempty"`

	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`

	Dt int64 `json:"dt"`

	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`

	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}
