package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// DownloadWeather is the function that initiates the download loop / retry
func DownloadWeather() {
	config := GetConfig("conf.json")

	for {
		err := downloadWeather()

		if err == nil {
			break
		}

		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Retrying in %v minutes...\n", config.Retry[0])
		time.Sleep(time.Duration(config.Retry[0]) * time.Minute)
	}

}

// downloadWeather fuction initiates the url check and download the data if possible
func downloadWeather() error {

	config := GetConfig("conf.json")

	StationId := config.StationId[0]
	OpenWeatherApi := config.OpenWeatherApi[0]
	WundergroundApi := config.WundergroundApi[0]

	if len(OpenWeatherApi) != 32 || ValidateAPICount(OpenWeatherApi) {
		OpenWeatherApi = "invalid"
	}

	if len(WundergroundApi) != 32 || ValidateAPICount(WundergroundApi) {
		WundergroundApi = "invalid"
	}

	if StationId == "" {
		StationId = "invalid"
	}

	if WundergroundApi != "invalid" && StationId != "invalid" && OpenWeatherApi != "invalid" {
		valid, _, weatherData := getWunderground()
		if !valid {
			valid, _, weatherData = getOpenWeather()
			if !valid {
				return fmt.Errorf("### Failed to collect weather from Wunderground & OpenWeather! ###")
			} else {
				SaveWeatherData(weatherData)
			}
		} else {
			SaveWeatherData(weatherData)
		}

	} else if WundergroundApi != "invalid" && StationId != "invalid" && OpenWeatherApi == "invalid" {
		valid, _, weatherData := getWunderground()
		if !valid {
			SaveWeatherData(weatherData)
		} else {
			return fmt.Errorf("### Failed to collect weather from Wunderground! ###")
		}
	} else if OpenWeatherApi != "invalid" {
		valid, _, weatherData := getOpenWeather()
		if !valid {
			return fmt.Errorf("### Failed to collect weather from OpenWeather! ###")
		} else {
			SaveWeatherData(weatherData)
		}
	} else {
		fmt.Println("No valid API configuration found")
	}
	return nil
}

func getOpenWeather() (bool, string, WeatherData) {
	var url, error string
	status := true

	config := GetConfig("conf.json")

	url = fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&lang=%s&units=metric&APPID=%s",
		config.DefaultCity[0], config.Language[0], config.OpenWeatherApi[0])

	response, err := http.Get(url)
	if err != nil {
		error = fmt.Sprintf("Error: Openweathermap making API request: %s\n", err)
		status = false
	}
	defer func() {
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}()

	var weatherData WeatherData

	if response != nil {

		if response.StatusCode != http.StatusOK {
			error = fmt.Sprintf("Openweathermap API Error: %s\n", response.Status)
			status = false
		}

		jsonData, err := io.ReadAll(response.Body)
		if err != nil {
			error = fmt.Sprintf("Error: Openweathermap reading response body: %s\n", err)
			status = false
		}

		var openweatherData Openweathermap

		if status {
			err = json.Unmarshal(jsonData, &openweatherData)
			if err != nil {
				error = fmt.Sprintf("Error: Openweathermap decoding JSON: %s\n", err)
				status = false
			}
		}

		weatherData.Observations = make([]Observation, 1)

		if status {
			TimeUtc := ConvertSeconds(int64(openweatherData.Dt))
			TimeLocal := strings.ReplaceAll(TimeUtc, "T", " ")
			TimeLocal = strings.ReplaceAll(TimeLocal, "Z", "")

			weatherData.Observations[0].ObsTimeUtc = ConvertTime(int64(openweatherData.Dt))
			weatherData.Observations[0].ObsTimeLocal = TimeLocal
			weatherData.Observations[0].Neighborhood = openweatherData.Name
			weatherData.Observations[0].Country = openweatherData.Sys.Country
			weatherData.Observations[0].SolarRadiation = 0.0
			weatherData.Observations[0].Lon = openweatherData.Coord.Lon
			weatherData.Observations[0].RealtimeFrequency = 0.0
			weatherData.Observations[0].Epoch = openweatherData.Dt
			weatherData.Observations[0].Lat = openweatherData.Coord.Lat
			weatherData.Observations[0].UV = 0.0
			weatherData.Observations[0].Winddir = 0
			weatherData.Observations[0].Humidity = openweatherData.Main.Humidity
			weatherData.Observations[0].QCStatus = 1
			weatherData.Observations[0].Metric.Temp = openweatherData.Main.Temp
			weatherData.Observations[0].Metric.HeatIndex = openweatherData.Main.Temp

			number := (openweatherData.Main.Temp - ((100 - openweatherData.Main.Humidity) / 7))
			_, err := fmt.Sscanf(fmt.Sprintf("%.2f", number), "%f", &number)
			if err != nil {
				fmt.Println("Error:", err)
			}

			weatherData.Observations[0].Metric.Dewpt = float64(number)
			weatherData.Observations[0].Metric.WindChill = 0.0
			weatherData.Observations[0].Metric.WindSpeed = openweatherData.Wind.Speed
			weatherData.Observations[0].Metric.WindGust = openweatherData.Wind.Gust
			weatherData.Observations[0].Metric.Pressure = openweatherData.Main.Pressure
			weatherData.Observations[0].Metric.PrecipRate = 0.0
			weatherData.Observations[0].Metric.PrecipTotal = 0.0
			weatherData.Observations[0].Description = openweatherData.Weather[0].Description
		}

		if error == "" {
			error = "OpenWeathermap"
		}
	} else {
		status = false
		error = "failed"
	}

	return status, error, weatherData
}

func getWunderground() (bool, string, WeatherData) {
	var url, error string
	status := true

	config := GetConfig("conf.json")

	url = fmt.Sprintf("https://api.weather.com/v2/pws/observations/current?stationId=%s&format=json&units=m&apiKey=%s&numericPrecision=decimal",
		config.StationId[0], config.WundergroundApi[0])

	response, err := http.Get(url)
	if err != nil {
		error = fmt.Sprintf("Error: WundergroundApi making API request: %s\n", err)
		status = false
	}
	defer func() {
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}()

	var weatherData WeatherData

	if response != nil {

		if response.StatusCode != http.StatusOK {
			error = fmt.Sprintf("WundergroundApi API Error: %s\n", response.Status)
			status = false
		}

		jsonData, err := io.ReadAll(response.Body)
		if err != nil {
			error = fmt.Sprintf("Error: WundergroundApi reading response body: %s\n", err)
			status = false
		}

		weatherData.Observations = make([]Observation, 1)
		weatherData.Observations[0].Description = "Wunderground"

		if status {
			err = json.Unmarshal([]byte(jsonData), &weatherData)
			if err != nil {
				error = fmt.Sprintf("Error: WundergroundApi decoding JSON: %s\n", err)
				status = false
			}
		}

		if weatherData.Observations[0].Humidity == 0 && weatherData.Observations[0].Metric.Dewpt == 0 {
			status = false
		}

		if error == "" {
			error = "Wunderground"
		}

	} else {
		status = false
		error = "failed"
	}

	return status, error, weatherData
}
