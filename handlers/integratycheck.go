package handlers

import (
	"database/sql"
	"encoding/json"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func SetupIntegraty() string {

	status := true
	msg := ""

	for status {
		// Json config file PATH
		pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		filePath := "/conf/conf.json"
		config := GetConfig()

		// Check if config file exist
		jsonData, err := os.ReadFile(pwd + filePath)
		if err != nil {
			msg += "\n>> Config File Missing! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Config File Present & accessible! Status OK"
			status = true
		}

		// Unmarshal JSON data into Config struct
		if err = json.Unmarshal(jsonData, &config); err != nil {
			msg += "\n>> Error unmarshaling Config! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Unmarshaling Config File! Status OK"
			status = true
		}

		// Check if the integrity is good (ignoring values)
		// You can add more specific checks based on your requirements
		if len(config.DB_HOST) == 1 &&
			len(config.DB_PORT) == 1 &&
			len(config.DB_USER) == 1 &&
			len(config.DB_PASS) == 1 &&
			len(config.DB_NAME) == 1 &&
			len(config.TB_NAME) == 1 &&
			len(config.OpenWeatherApi) == 1 &&
			len(config.WundergroundApi) == 1 &&
			len(config.StationId) == 1 &&
			len(config.WebPort) == 1 &&
			len(config.Language) == 1 &&
			len(config.DefaultCity) == 1 &&
			len(config.PrivKeyPATH) == 1 &&
			len(config.CertPemPATH) == 1 &&
			len(config.Retry) == 1 &&
			len(config.Interval) == 1 {
			msg += "\n>> Reading Config integrity! Status OK"
			status = true
		} else {
			msg += "\n>> Reading Config integrity! Status Failed"
			status = false
			break
		}

		// Reading DB config values
		if config.DB_HOST[0] == "" {
			msg += "\n>> Reading DB_HOST isEmpty! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Reading DB_HOST hasValue! Status OK (" + config.DB_HOST[0] + ")"
			status = true
		}
		if config.DB_PORT[0] == "" {
			msg += "\n>> Reading DB_PORT isEmpty! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Reading DB_PORT hasValue! Status OK (" + config.DB_PORT[0] + ")"
			status = true
		}
		if config.DB_USER[0] == "" {
			msg += "\n>> Reading DB_USER isEmpty! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Reading DB_USER hasValue! Status OK (" + config.DB_USER[0] + ")"
			status = true
		}
		if config.DB_PASS[0] == "" {
			msg += "\n>> Reading DB_PASS isEmpty! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Reading DB_PASS hasValue! Status OK (" + config.DB_PASS[0] + ")"
			status = true
		}
		if config.DB_NAME[0] == "" {
			msg += "\n>> Reading DB_NAME isEmpty! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Reading DB_NAME hasValue! Status OK (" + config.DB_NAME[0] + ")"
			status = true
		}
		if config.TB_NAME[0] == "" {
			msg += "\n>> Reading TB_NAME isEmpty! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Reading TB_NAME hasValue! Status OK (" + config.TB_NAME[0] + ")"
			status = true
		}

		db, _ := GetDBConnection()
		err = db.Ping()
		if err != nil {
			msg += "\n>> Checking Database Connection! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Checking Database Connection! Status OK"
			status = false
		}

		tableName := config.TB_NAME[0]
		expectedColumns := []string{
			"id", "obstimeutc", "obstimelocal", "neighborhood", "country", "solarradiation",
			"lon", "realtimefrequency", "epoch", "lat", "uv", "winddir", "humidity", "qcstatus",
			"temp", "heatindex", "dewpt", "windchill", "windspeed", "windgust", "pressure", "preciprate",
			"preciptotal", "created_at", "freetext",
		}

		defer db.Close()
		exists, err := columnsExist(db, tableName, expectedColumns)
		if err != nil {
			msg += "\n>> Reading Tables & Columns! Status Failed"
			status = false
			break
		}

		if exists {
			msg += "\n>> All Table & Columns exists! Status OK"
			status = true
		} else {
			msg += "\n>> Reading Tables & Columns! Status Failed"
			status = false
			break
		}

		if isEmpty(config.WebPort[0]) || !isNumeric(config.WebPort[0]) {
			msg += "\n>> WebPort has no Valid Port! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Reading WebPort hasValue! Status OK (" + config.WebPort[0] + ")"
			status = true
		}

		timeout := time.Second
		host := GetOutboundIP()
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host.String(), config.WebPort[0]), timeout)
		if err != nil {
			msg += "\n>> Tested WebPort isOpen! Status OK"
			status = true
		} else {
			defer conn.Close()
			msg += "\n>> WebPort has no Valid Port! Status Failed"
			status = false
			break
		}

		if config.OpenWeatherApi[0] == "" && config.WundergroundApi[0] == "" {
			msg += "\n>> All API key configured isEmpty! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> Found one or more API key! Status OK"
			status = true
		}
		if config.OpenWeatherApi[0] != "" && len(config.OpenWeatherApi[0]) == 32 {
			msg += "\n>> OpenWeatherApi API key length! Status OK  (" + config.OpenWeatherApi[0] + ")"
			status = true
		} else {
			msg += "\n>> OpenWeatherApi API key length! Status Failed"
			status = false
			break
		}
		if config.WundergroundApi[0] != "" && len(config.WundergroundApi[0]) == 32 {
			msg += "\n>> WundergroundApi API key length! Status OK  (" + config.WundergroundApi[0] + ")"
			status = true
		} else {
			msg += "\n>> WundergroundApi API key length! Status Failed"
			status = false
			break
		}
		if !isEmptyInt(config.Interval[0]) {
			msg += "\n>> Interval has IntegerValue! Status OK (" + strconv.Itoa(config.Interval[0]) + " Minutes)"
			status = true
		} else {
			msg += "\n>> Interval NO IntegerValue! Status Failed"
			status = false
			break
		}
		if isEmpty(config.DefaultCity[0]) {
			msg += "\n>> DefaultCity is NOT configured! Status Failed"
			status = false
			break
		} else {
			msg += "\n>> DefaultCity has StringValue! Status OK (" + config.DefaultCity[0] + ")"
			status = true
		}

		if isEmpty(config.CertPemPATH[0]) {
			msg += "\n>> CertPemPATH is NOT configured SSL unavailable! Status Warning"
			status = true
		} else {
			msg += "\n>> CertPemPATH has StringValue! Status OK"
			status = true
			if strings.HasSuffix(config.CertPemPATH[0], ".pem") {
				msg += "\n>> CertPemPATH seems valid (pem)! Status OK (" + config.CertPemPATH[0] + ")"
				status = true
			} else {
				msg += "\n>> CertPemPATH configured BUT invalid (pem)! Status Failed"
				status = false
				break
			}
		}

		if isEmpty(config.PrivKeyPATH[0]) {
			msg += "\n>> PrivKeyPATH is NOT configured SSL unavailable! Status Warning"
			status = true
			break
		} else {
			msg += "\n>> PrivKeyPATH has StringValue! Status OK"
			status = true
			if strings.HasSuffix(config.PrivKeyPATH[0], ".pem") {
				msg += "\n>> PrivKeyPATH seems valid (pem)! Status OK (" + config.PrivKeyPATH[0] + ")"
				status = true
			} else {
				msg += "\n>> PrivKeyPATH configured BUT invalid (pem)! Status Failed"
				status = false
				break
			}
		}

		if !isEmptyInt(config.Retry[0]) {
			msg += "\n>> Retry has IntegerValue! Status OK (" + strconv.Itoa(config.Retry[0]) + " Minutes)"
			status = true
		} else {
			msg += "\n>> Retry NO IntegerValue! Status Failed"
			status = false
			break
		}

		// End of checks
		msg += "\n"
		status = false
	}

	return msg
}

func isEmpty(s string) bool {
	return s == ""
}

func isEmptyInt(i int) bool {
	return i == 0
}

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func columnsExist(db *sql.DB, tableName string, expectedColumns []string) (bool, error) {
	query := `
		SELECT column_name
		FROM information_schema.columns
		WHERE table_name = $1
	`
	rows, err := db.Query(query, tableName)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	// Collect existing column names in a slice
	existingColumns := make([]string, 0)
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			return false, err
		}
		existingColumns = append(existingColumns, columnName)
	}

	// Check if all expected columns exist in the table
	for _, col := range expectedColumns {
		found := false
		for _, existingCol := range existingColumns {
			if col == existingCol {
				found = true
				break
			}
		}
		if !found {
			return false, nil
		}
	}

	return true, nil
}
