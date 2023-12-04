package handlers

import "time"

// Config File Structure
type Config struct {
	DB_HOST         []string `json:"DB_HOST"`
	DB_PORT         []string `json:"DB_PORT"`
	DB_USER         []string `json:"DB_USER"`
	DB_PASS         []string `json:"DB_PASS"`
	DB_NAME         []string `json:"DB_NAME"`
	TB_NAME         []string `json:"TB_NAME"`
	OpenWeatherApi  []string `json:"OpenWeatherApi"`
	StationValid    []bool   `json:"StationValid"`
	WundergroundApi []string `json:"WundergroundApi"`
	StationId       []string `json:"StationId"`
	WebPort         []string `json:"WebPort"`
	Language        []string `json:"Language"`
	DefaultCity     []string `json:"DefaultCity"`
}

// Metric represents the nested "metric" object in the JSON structure
type Metric struct {
	Temp        float64 `json:"temp"`
	HeatIndex   float64 `json:"heatIndex"`
	Dewpt       float64 `json:"dewpt"`
	WindChill   float64 `json:"windChill"`
	WindSpeed   float64 `json:"windSpeed"`
	WindGust    float64 `json:"windGust"`
	Pressure    float64 `json:"pressure"`
	PrecipRate  float64 `json:"precipRate"`
	PrecipTotal float64 `json:"precipTotal"`
	Elev        float64 `json:"elev"`
}

// Observation represents each observation in the "observations" array
type Observation struct {
	StationID         string    `json:"stationID"`
	ObsTimeUtc        time.Time `json:"obsTimeUtc"`
	ObsTimeLocal      string    `json:"obsTimeLocal"`
	Neighborhood      string    `json:"neighborhood"`
	SoftwareType      string    `json:"softwareType"`
	Country           string    `json:"country"`
	SolarRadiation    float64   `json:"solarRadiation"`
	Lon               float64   `json:"lon"`
	RealtimeFrequency int       `json:"realtimeFrequency"`
	Epoch             int       `json:"epoch"`
	Lat               float64   `json:"lat"`
	UV                float64   `json:"uv"`
	Winddir           float64   `json:"winddir"`
	Humidity          float64   `json:"humidity"`
	QCStatus          float64   `json:"qcStatus"`
	Metric            Metric    `json:"metric"`
	Description       string    `json:"description"`
}

// WeatherData represents the overall JSON structure
type WeatherData struct {
	Observations []Observation `json:"observations"`
}

type Openweathermap struct {
	Coord      Coord   `json:"coord"`
	Weather    Weather `json:"weather"`
	Base       string  `json:"base"`
	Main       Main    `json:"main"`
	Visibility int     `json:"visibility"`
	Wind       Wind    `json:"wind"`
	Rain       Rain    `json:"rain"`
	Clouds     Clouds  `json:"clouds"`
	Dt         int     `json:"dt"`
	Sys        Sys     `json:"sys"`
	Timezone   int32   `json:"timezone"`
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Cod        int     `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Rain struct {
	OneHour float64 `json:"1h"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}
