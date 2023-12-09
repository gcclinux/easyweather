package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func LaunchWeb() {
	var address = fmt.Sprintf("%s:%s", GetOutboundIP(), GetConfig().WebPort[0])

	// Define the endpoint
	http.HandleFunc("/json", JsonHandler)
	http.HandleFunc("/", HomeHandler)

	// Adding additional required directories
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	fmt.Printf("Listening on http://%s\n", address)

	http.ListenAndServe(address, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	weatherData := GetWeatherData()

	temperature, humidity, preHumidity, preDewpt, prePressure, preWindspeed, windspeed, dewpt, pressure := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	var Location, Minwhen, mini, date, time = "", "", "", "", ""
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
		Temp:         temperature,
		Humidity:     humidity,
		Neighborhood: Location,
		PreLow:       mini,
		PreWhen:      Minwhen,
		Date:         date,
		PreHumidity:  preHumidity,
		PreDewpt:     preDewpt,
		PrePressure:  prePressure,
		Dewpt:        dewpt,
		PreWindspeed: preWindspeed,
		WindSpeed:    windspeed,
		Pressure:     pressure,
		Time:         time,
	}

	tmpl, err := template.ParseFiles("html/home.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, ParseWeather)
}
