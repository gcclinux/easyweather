package handlers

import (
	"fmt"
	"log"
)

func GetWeatherData() []WeatherData {

	db, err := GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Query the database
	rows, err := db.Query("SELECT obstimelocal, neighborhood, winddir, humidity, temp, dewpt, windchill, windspeed, pressure FROM ecowitt_weather WHERE ecowitt_weather.obstimelocal >= CURRENT_DATE AND ecowitt_weather.obstimelocal < (CURRENT_DATE + '1 day'::interval) ORDER BY ecowitt_weather.id ASC; ")
	if err != nil {
		log.Fatal(err)
		fmt.Println("GetWeatherData FAILED")
	} else {
		fmt.Println("GetWeatherData OK")
	}
	defer rows.Close()

	// Slice to store WeatherData objects
	var weatherDataSlice []WeatherData

	// Loop through the result set
	for rows.Next() {
		var weatherData WeatherData
		err := rows.Scan(&weatherData.Observations[0].ObsTimeLocal, &weatherData.Observations[0].Neighborhood, &weatherData.Observations[0].Winddir, &weatherData.Observations[0].Humidity, &weatherData.Observations[0].Metric.Temp, &weatherData.Observations[0].Metric.Dewpt, &weatherData.Observations[0].Metric.WindChill, &weatherData.Observations[0].Metric.WindSpeed, &weatherData.Observations[0].Metric.Pressure)
		if err != nil {
			log.Fatal(err)
		}

		// Append WeatherData to the slice
		weatherDataSlice = append(weatherDataSlice, weatherData)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return weatherDataSlice
}
