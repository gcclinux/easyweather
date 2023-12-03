package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// GetConfig() returns all the conf.json values
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
