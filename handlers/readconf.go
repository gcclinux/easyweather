package handlers

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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
		WundergroundApi []string `json:"WundergroundApi"`
		StationId       []string `json:"StationId"`
		WebPort         []string `json:"WebPort"`
		Language        []string `json:"Language"`
		DefaultCity     []string `json:"DefaultCity"`
		Interval        []int    `json:"Interval"`
		Retry 			[]int	 `json:"Retry"`
	    PrivKeyPATH     []string `json:"PrivKeyPATH"`
	    CertPemPATH     []string `json:"CertPemPATH"`
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

func GetSSLName() string {

	pemFilePath := GetConfig().CertPemPATH[0]
	value := ""

	// Read the PEM file
	pemData, err := os.ReadFile(pemFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the PEM block
	block, _ := pem.Decode(pemData)
	if block == nil {
		log.Fatal("Failed to parse PEM block")
	}

	// Parse the DER-encoded X.509 certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	certName := fmt.Sprintf("%s", cert.Subject)
	parts := strings.Split(certName, "=")

	if len(parts) == 2 {
		value = parts[1]
	} else {
		fmt.Println("Invalid cert.pem format")
	}
	return value
}
