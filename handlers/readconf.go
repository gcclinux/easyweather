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
	    DB_HOST        []string `json:"DB_HOST"`
	    DB_PORT        []string `json:"DB_PORT"`
	    DB_USER        []string `json:"DB_USER"`
	    DB_PASS        []string `json:"DB_PASS"`
	    DB_NAME        []string `json:"DB_NAME"`
	    TB_NAME        []string `json:"TB_NAME"`
	    OpenWeatherApi []string `json:"OpenWeatherApi"`
	    StationValid   []bool   `json:"StationValid"`
	    WundergroupApi []string `json:"WundergroupApi"`
	    StationId      []string `json:"StationId"`
	    AdminPort      []string `json:"AdminPort"`
	}
*/
func GetConfig() Config {
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := "/conf/conf.json"
	file, _ := os.Open(pwd + filePath)

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
