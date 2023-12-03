package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {

	configuration := GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.DB_HOST[0], configuration.DB_PORT[0], configuration.DB_USER[0], configuration.DB_PASS[0], configuration.DB_NAME[0])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db, nil
}
