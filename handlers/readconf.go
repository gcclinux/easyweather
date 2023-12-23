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
GetConfig() returns all the a json file variables
*/
func GetConfig(name string) Config {
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := "/conf/" + name

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

	pemFilePath := GetConfig("conf.json").CertPemPATH[0]
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

	certName := cert.Subject.String()
	parts := strings.Split(certName, "=")

	if len(parts) == 2 {
		value = parts[1]
	} else {
		fmt.Println("Invalid cert.pem format")
	}
	return value
}
