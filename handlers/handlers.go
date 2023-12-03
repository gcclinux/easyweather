package handlers

import "time"

// Config File Structure
type Config struct {
	DB_HOST        []string `json:"DB_HOST"`
	DB_PORT        []string `json:"DB_PORT"`
	DB_USER        []string `json:"DB_USER"`
	DB_PASS        []string `json:"DB_PASS"`
	DB_NAME        []string `json:"DB_NAME"`
	TB_NAME        []string `json:"TB_NAME"`
	OpenWeatherApi []string `json:"OpenWeatherApi"`
	StationValid   []string `json:"StationValid"`
	WundergroupApi []string `json:"WundergroupApi"`
	StationId      []string `json:"StationId"`
	AdminPort      []string `json:"AdminPort"`
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
}

// WeatherData represents the overall JSON structure
type WeatherData struct {
	Observations []Observation `json:"observations"`
}
