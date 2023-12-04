package main

import (
	"weatherCollect/handlers"
)

func main() {

	handlers.GetDBConnection()
	handlers.DownloadWeather()

}
