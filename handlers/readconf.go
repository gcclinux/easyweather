package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

/*
GetConfig() returns all the conf.json values

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
	Interval        []int    `json:"Interval"`
	}
*/
func GetConfig() Config {
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := "/conf/conf.json"
	file, err := os.Open(pwd + filePath)
	if err != nil {
		fmt.Println("file error:", err)
		os.Exit(3)
	}

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error in decoding conf.json:", err)
	}
	return configuration
}
