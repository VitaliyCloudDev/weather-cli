package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetGeo(api string, city string, limit string) []byte {
	var geoJson []byte = get_request(api, city, limit)
	geo, err := format_get(geoJson)
	if err != nil {
		fmt.Println(err)
	}
	data := get_data(geo, api)
	return data
}

func get_request(api string, city string, limit string) []byte {
	raw_url := fmt.Sprintf(
		"http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=%s",
		url.QueryEscape(city),
		url.QueryEscape(limit),
	)
	return processraw_url(raw_url, api)
}

func get_data(geo []CityInfo, api string) []byte {
	raw_url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=metric",
		geo[0].Lat,
		geo[0].Lon,
	)
	raw_json := processraw_url(raw_url, api)
	return raw_json
}

func processraw_url(raw_url string, api string) []byte {
	url_request := fmt.Sprintf("%s&appid=%s", raw_url, url.QueryEscape(api))
	response, err := http.Get(url_request)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса: ", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении результата: ", err)
	}
	return body
}
