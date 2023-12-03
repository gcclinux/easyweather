package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DownloadWeather() {

	config := GetConfig()

	StationId := config.StationId[0]
	OpenWeatherApi := config.OpenWeatherApi[0]
	WundergroundApi := config.WundergroundApi[0]
	StationValid := false

	if len(OpenWeatherApi) != 32 || ValidateAPICount(OpenWeatherApi) {
		OpenWeatherApi = "invalid"
	}

	if config.StationValid[0] {
		StationValid = true
	}

	if len(WundergroundApi) != 32 || ValidateAPICount(WundergroundApi) {
		WundergroundApi = "invalid"
	}

	if StationId == "" {
		StationId = "invalid"
	}

	if StationValid && WundergroundApi != "invalid" && StationId != "invalid" && OpenWeatherApi != "invalid" {
		valid, error, _ := getWunderground()
		if !valid {
			valid, _ = getOpenWeather()
			if !valid {
				fmt.Println("### Failed to collect weather! ###")
				fmt.Println(error)
			} else {
				fmt.Println("Ready to upload")
			}
		} else {
			fmt.Println("Ready to upload")
		}
	} else if OpenWeatherApi != "invalid" {
		fmt.Println("OpenWeatherApi: ", OpenWeatherApi)
	} else {
		fmt.Println("No valid API configuration found")
	}

}

func getOpenWeather() (bool, WeatherData) {
	var weatherData WeatherData
	return false, weatherData
}

func getWunderground() (bool, string, WeatherData) {
	var url, error string
	status := true

	config := GetConfig()

	url = fmt.Sprintf("https://api.weather.com/v2/pws/observations/current?stationId=%s&format=json&units=m&apiKey=%s&numericPrecision=decimal",
		config.StationId[0], config.WundergroundApi[0])

	response, err := http.Get(url)
	if err != nil {
		error = fmt.Sprintf("Error: WundergroundApi making API request: %s\n", err)
		status = false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		error = fmt.Sprintf("WundergroundApi API Error: %s\n", response.Status)
		status = false
	}

	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		error = fmt.Sprintf("Error: WundergroundApi reading response body: %s\n", err)
		status = false
	}

	var weatherData WeatherData
	if status {
		err = json.Unmarshal([]byte(jsonData), &weatherData)
		if err != nil {
			error = fmt.Sprintf("Error: WundergroundApi decoding JSON: %s\n", err)
			status = false
		}
	}

	return status, error, weatherData
}