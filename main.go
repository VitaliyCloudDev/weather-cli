package main

import (
	"flag"

	"weather-cli.go/internal/weather"
)

var city string
var api_key string

func main() {
	init_flags()
	geo := weather.GetGeo(api_key, city, "1")
	weather.ShowInfo(geo)
}
func init_flags() {
	flag.StringVar(&api_key, "api", "", "Api key from Openweather.")
	flag.StringVar(&city, "city", "Moscow", "City to get data.")
	flag.Parse()
}
