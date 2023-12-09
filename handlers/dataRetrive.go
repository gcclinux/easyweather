package handlers

import (
	"fmt"
	"log"
)

func GetWeatherData() []WeatherStruct {

	db, err := GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Query the database
	rows, err := db.Query("SELECT obstimelocal, neighborhood, winddir, humidity, temp, dewpt, windchill, windspeed, pressure FROM ecowitt_weather WHERE ecowitt_weather.obstimelocal >= CURRENT_DATE AND ecowitt_weather.obstimelocal < (CURRENT_DATE + '1 day'::interval) ORDER BY ecowitt_weather.id ASC; ")
	if err != nil {
		log.Fatal(err)
		fmt.Println("GetWeatherData FAILED")
	}

	defer rows.Close()

	// Slice to store WeatherData objects
	var weatherDataSlice []WeatherStruct

	// Loop through the result set
	for rows.Next() {
		var weatherData WeatherStruct
		err := rows.Scan(&weatherData.Obstimelocal, &weatherData.Neighborhood, &weatherData.Winddir, &weatherData.Humidity, &weatherData.Temp, &weatherData.Dewpt, &weatherData.WindChill, &weatherData.WindSpeed, &weatherData.Pressure)
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
