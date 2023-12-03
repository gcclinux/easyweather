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
	WundergroupApi := config.WundergroupApi[0]
	StationValid := false

	if len(OpenWeatherApi) != 32 || ValidateAPICount(OpenWeatherApi) {
		OpenWeatherApi = "invalid"
	}

	if config.StationValid[0] {
		StationValid = true
	}

	if len(WundergroupApi) != 32 || ValidateAPICount(WundergroupApi) {
		WundergroupApi = "invalid"
	}

	if StationId == "" {
		StationId = "invalid"
	}

	if StationValid && WundergroupApi != "invalid" && StationId != "invalid" && OpenWeatherApi != "invalid" {
		valid, _ := getWundergroup()
		if !valid {
			valid, _ = getOpenWeather()
			if !valid {
				fmt.Println("Failed to collect weather, check your internet connection!")
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
	panic("unimplemented")
}

func getWundergroup() (bool, WeatherData) {
	var url string
	status := true
	config := GetConfig()

	url = fmt.Sprintf("https://api.weather.com/v2/pws/observations/current?stationId=%s&format=json&units=m&apiKey=%s&numericPrecision=decimal",
		config.StationId[0], config.WundergroupApi[0])

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making API request: %s\n", err)
		status = false
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", response.Status)
		status = false
	}

	jsonData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		status = false
	}

	var weatherData WeatherData
	err = json.Unmarshal([]byte(jsonData), &weatherData)
	if err != nil {
		fmt.Printf("Error decoding JSON: %s\n", err)
		status = false
	}

	return status, weatherData
}
