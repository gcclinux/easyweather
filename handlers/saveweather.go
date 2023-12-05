package handlers

import (
	"log"
)

func SaveWeatherData(weatherData WeatherData) {

	db, err := GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("SaveWeatherData: %+v", weatherData)

	obsTimeUtc := weatherData.Observations[0].ObsTimeUtc.String()[:19]
	obsTimeLocal := weatherData.Observations[0].ObsTimeLocal
	neighborhood := weatherData.Observations[0].Neighborhood
	country := weatherData.Observations[0].Country
	solarRadiation := weatherData.Observations[0].SolarRadiation
	lon := weatherData.Observations[0].Lon
	realtimeFrequency := weatherData.Observations[0].RealtimeFrequency
	epoch := weatherData.Observations[0].Epoch
	lat := weatherData.Observations[0].Lat
	uv := weatherData.Observations[0].UV
	winddir := weatherData.Observations[0].Winddir
	humidity := weatherData.Observations[0].Humidity
	qcStatus := weatherData.Observations[0].QCStatus
	temp := weatherData.Observations[0].Metric.Temp
	heatIndex := weatherData.Observations[0].Metric.HeatIndex
	dewpt := weatherData.Observations[0].Metric.Dewpt
	windChill := weatherData.Observations[0].Metric.WindChill
	windSpeed := weatherData.Observations[0].Metric.WindSpeed
	windGust := weatherData.Observations[0].Metric.WindGust
	pressure := weatherData.Observations[0].Metric.Pressure
	precipRate := weatherData.Observations[0].Metric.PrecipRate
	precipTotal := weatherData.Observations[0].Metric.PrecipTotal
	description := weatherData.Observations[0].Description

	defer db.Close()

	_, err = db.Exec("INSERT INTO public.ecowitt_weather(obsTimeUtc,obsTimeLocal,neighborhood,country,solarRadiation,lon,realtimeFrequency,epoch,lat,uv,winddir,humidity,qcStatus,temp,heatIndex,dewpt,windChill,windSpeed,windGust,pressure,precipRate,precipTotal,freetext) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)",
		obsTimeUtc, obsTimeLocal, neighborhood, country, solarRadiation, lon, realtimeFrequency, epoch, lat, uv, winddir, humidity, qcStatus, temp, heatIndex, dewpt, windChill, windSpeed, windGust, pressure, precipRate, precipTotal, description)
	if err != nil {
		log.Fatal("db.Exec ", err)
	}

	// fmt.Println("obsTimeUtc: ", obsTimeUtc)
	// fmt.Println("obsTimeLocal: ", obsTimeLocal)
	// fmt.Println("neighborhood: ", neighborhood)
	// fmt.Println("country: ", country)
	// fmt.Println("solarRadiation: ", solarRadiation)
	// fmt.Println("lon: ", lon)
	// fmt.Println("realtimeFrequency: ", realtimeFrequency)
	// fmt.Println("epoch: ", epoch)
	// fmt.Println("lat: ", lat)
	// fmt.Println("uv: ", uv)
	// fmt.Println("winddir: ", winddir)
	// fmt.Println("humidity: ", humidity)
	// fmt.Println("qcStatus: ", qcStatus)
	// fmt.Println("temp: ", temp)
	// fmt.Println("heatIndex: ", heatIndex)
	// fmt.Println("dewpt: ", dewpt)
	// fmt.Println("windChill: ", windChill)
	// fmt.Println("windSpeed: ", windSpeed)
	// fmt.Println("windGust: ", windGust)
	// fmt.Println("pressure: ", pressure)
	// fmt.Println("precipRate: ", precipRate)
	// fmt.Println("precipTotal: ", precipTotal)
	// fmt.Println("description: ", description)
	// fmt.Println("")
}
