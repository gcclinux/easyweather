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

	var MaxT, temperature, humidity, preHumidity, preDewpt, prePressure, preWindspeed, windspeed, dewpt, pressure, precipTotal float64
	var Location, Minwhen, mini, maxi, date, time, remoteDate string
	MinT := math.Inf(0)

	for _, data := range weatherData {
		temperature = data.Temp
		humidity = data.Humidity
		Location = data.Neighborhood
		date = formatDate(data.Obstimelocal[:10])
		time = data.Obstimelocal[11:19]
		windspeed = data.WindSpeed
		dewpt = data.Dewpt
		pressure = data.Pressure
		precipTotal = data.PrecipTotal
		if GetConfig("conf.json").AdjustTime[0] {
			remoteDate = AdjustTimeTime(data.Obstimelocal, GetConfig("conf.json").TimeZone[0])[11:19]
		} else {
			remoteDate = time
		}

		if data.Temp < MinT {
			MinT = data.Temp
		}

		if data.Temp > MaxT {
			MaxT = data.Temp
		}
	}

	for _, dt := range weatherData {
		if MinT == dt.Temp {
			if GetConfig("conf.json").AdjustTime[0] {
				Minwhen = AdjustTimeTime(dt.Obstimelocal, GetConfig("conf.json").TimeZone[0])[11:19]
			} else {
				Minwhen = dt.Obstimelocal[11:19]
			}
			preHumidity = dt.Humidity
			preDewpt = dt.Dewpt
			prePressure = dt.Pressure
			preWindspeed = dt.WindSpeed
		}
	}

	// Converting Float64 to String & counting digits
	miniString := strconv.FormatFloat(MinT, 'f', -1, 64)
	if !strings.Contains(miniString, ".") {
		mini = fmt.Sprintf("%s.0", miniString)
	} else {
		mini = miniString
	}

	maxiString := strconv.FormatFloat(MaxT, 'f', -1, 64)
	if !strings.Contains(maxiString, ".") {
		maxi = fmt.Sprintf("%s.0", maxiString)
	} else {
		maxi = maxiString
	}

	ParseWeather := WeatherStruct{
		Temp:          temperature,
		Humidity:      humidity,
		Neighborhood:  Location,
		RemoteDate:    remoteDate,
		PreLow:        mini,
		Max:           maxi,
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
		LangSpeed:     language.WINDSPEED[0],
		LangRainTotal: language.TOTALRAIN[0],
		LangPressure:  language.PRESSURE[0],
		LangLowest:    language.LOWEST[0],
		LangHighest:   language.HIGHEST[0],
		LangTime:      language.TIME[0],
		LangDate:      language.DATE[0],
	}

	tmpl, err := template.ParseFiles("html/home.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, ParseWeather)
}
