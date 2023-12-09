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

	for i := range weatherData {
		TimeLocal := weatherData[i].Observations[0].ObsTimeLocal[11:19]
		weatherData[i].Observations[0].ObsTimeLocal = TimeLocal
		if weatherData[i].Observations[0].Metric.Temp < MinT {
			MinT = weatherData[i].Observations[0].Metric.Temp
		}
	}

	for dt := range weatherData {
		if MinT == weatherData[dt].Observations[0].Metric.Temp {
			weatherData[dt].Observations[0].PreWhen = weatherData[dt].Observations[0].ObsTimeLocal
			weatherData[dt].Observations[0].PreHumidity = weatherData[dt].Observations[0].Humidity
			weatherData[dt].Observations[0].PreDewpt = weatherData[dt].Observations[0].Metric.Dewpt
			weatherData[dt].Observations[0].PrePressure = weatherData[dt].Observations[0].Metric.Pressure
			weatherData[dt].Observations[0].PrePressure = weatherData[dt].Observations[0].Metric.WindSpeed
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
