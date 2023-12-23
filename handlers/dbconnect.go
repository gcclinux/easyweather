package handlers

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {

	configuration := GetConfig("conf.json")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.DB_HOST[0], configuration.DB_PORT[0], configuration.DB_USER[0], configuration.DB_PASS[0], configuration.DB_NAME[0])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error GetDBConnection() sql.Open:", err)
		os.Exit(3)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error GetDBConnection() sql.Ping:", err)
		os.Exit(3)
	}
	return db, nil
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Error GetOutboundIP() conn:", err)
		os.Exit(3)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
