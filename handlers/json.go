package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	weatherData := GetWeatherData()
	MinT := math.Inf(0)
	var TimeLocal string

	for i := range weatherData {
		if GetConfig("conf.json").AdjustTime[0] {
			TimeLocal = AdjustTimeTime(weatherData[i].Obstimelocal, GetConfig("conf.json").TimeZone[0])[11:19]
		} else {
			TimeLocal = weatherData[i].Obstimelocal[11:19]
		}

		weatherData[i].Obstimelocal = TimeLocal
		if weatherData[i].Temp < MinT {
			MinT = weatherData[i].Temp
		}
	}

	for dt := range weatherData {
		if MinT == weatherData[dt].Temp {
			weatherData[dt].PreWhen = weatherData[dt].Obstimelocal
			weatherData[dt].PreHumidity = weatherData[dt].Humidity
			weatherData[dt].PreDewpt = weatherData[dt].Dewpt
			weatherData[dt].PrePressure = weatherData[dt].Pressure
			weatherData[dt].PreWindspeed = weatherData[dt].WindSpeed
		}
	}

	jsonWeather, err := json.Marshal(weatherData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response
	w.Write(jsonWeather)
}
