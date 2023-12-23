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

	config := GetConfig("conf.json")
	query := fmt.Sprintf("SELECT mt.obstimelocal, mt.neighborhood, mt.winddir, mt.humidity, mt.temp, mt.dewpt, mt.windchill, mt.windspeed, mt.pressure, mt.preciptotal FROM %s AS mt WHERE mt.neighborhood = '%s' AND mt.obstimelocal >= CURRENT_DATE AND mt.obstimelocal < (CURRENT_DATE + '1 day'::interval) ORDER BY mt.id ASC;", config.TB_NAME[0], config.DefaultCity[0])

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("GetWeatherData FAILED")
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice to store WeatherData objects
	var weatherDataSlice []WeatherStruct

	// Loop through the result set
	for rows.Next() {
		var weatherData WeatherStruct
		err := rows.Scan(&weatherData.Obstimelocal, &weatherData.Neighborhood, &weatherData.Winddir, &weatherData.Humidity, &weatherData.Temp, &weatherData.Dewpt, &weatherData.WindChill, &weatherData.WindSpeed, &weatherData.Pressure, &weatherData.PrecipTotal)
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
