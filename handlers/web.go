package handlers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func LaunchWeb() {
	var address = fmt.Sprintf("%s:%s", GetOutboundIP(), GetConfig("conf.json").WebPort[0])

	// Define the endpoint
	http.HandleFunc("/json", JsonHandler)
	http.HandleFunc("/", HomeHandler)

	// Adding additional required directories
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	if isEmpty(GetConfig("conf.json").PrivKeyPATH[0]) || isEmpty(GetConfig("conf.json").CertPemPATH[0]) {
		fmt.Printf("Listening on http://%s\n", address)
		http.ListenAndServe(address, nil)
	} else {
		fmt.Printf("Listening on https://%s:%s\n", GetSSLName(), GetConfig("conf.json").WebPort[0])
		err := http.ListenAndServeTLS(fmt.Sprintf("%v", address), GetConfig("conf.json").CertPemPATH[0], GetConfig("conf.json").PrivKeyPATH[0], nil)
		if err != nil {
			log.Println(err)
		}

	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	weatherData := GetWeatherData()
	language := getLanguage("language.json")

	var temperature, humidity, preHumidity, preDewpt, prePressure, preWindspeed, windspeed, dewpt, pressure, precipTotal float64
	var Location, Minwhen, mini, date, time string
	MinT := math.Inf(0)

	for _, data := range weatherData {
		temperature = data.Temp
		humidity = data.Humidity
		Location = data.Neighborhood
		date = data.Obstimelocal[:10]
		time = data.Obstimelocal[11:19]
		windspeed = data.WindSpeed
		dewpt = data.Dewpt
		pressure = data.Pressure
		precipTotal = data.PrecipTotal
		if data.Temp < MinT {
			MinT = data.Temp
		}
	}

	for _, dt := range weatherData {
		if MinT == dt.Temp {
			Minwhen = dt.Obstimelocal[11:19]
			preHumidity = dt.Humidity
			preDewpt = dt.Dewpt
			prePressure = dt.Pressure
			preWindspeed = dt.WindSpeed
		}
	}

	// Converting Float64 to String & counting digits
	numberString := strconv.FormatFloat(MinT, 'f', -1, 64)
	if !strings.Contains(numberString, ".") {
		mini = fmt.Sprintf("%s.0", numberString)
	} else {
		mini = numberString
	}

	ParseWeather := WeatherStruct{
		Temp:          temperature,
		Humidity:      humidity,
		Neighborhood:  Location,
		PreLow:        mini,
		PreWhen:       Minwhen,
		Date:          date,
		PreHumidity:   preHumidity,
		PreDewpt:      preDewpt,
		PrePressure:   prePressure,
		Dewpt:         dewpt,
		PreWindspeed:  preWindspeed,
		WindSpeed:     windspeed,
		Pressure:      pressure,
		Time:          time,
		PrecipTotal:   precipTotal,
		LangTemp:      language.TEMPERATURE[0],
		LangHum:       language.HUMIDITY[0],
		LangDew:       language.DEWPOINT[0],
		LangRainTotal: language.TOTALRAIN[0],
		LangPressure:  language.PRESSURE[0],
	}

	tmpl, err := template.ParseFiles("html/home.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, ParseWeather)
}
