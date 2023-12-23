package handlers

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// CollectLoop is the function that collect the data periodically based on the interval setting
func CollectLoop() {

	config := GetConfig("conf.json")

	// Create a channel to signal the goroutine to stop
	stopChan := make(chan struct{})

	// Run the command every 10 minutes
	ticker := time.NewTicker(time.Duration(config.Interval[0]) * time.Minute)

	// Run the command immediately upon startup
	DownloadWeather()

	// Use a WaitGroup to wait for the goroutine to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Define a goroutine to run the command at the specified interval
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				DownloadWeather()
			case <-stopChan:
				// Terminate the goroutine when the stop signal is received
				return
			}
		}
	}()

	// Handle OS signals to stop the goroutine on interrupt
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		close(stopChan)
	}()

	// Wait for the goroutine to finish before exiting
	wg.Wait()

	// Optionally perform cleanup before exiting
	fmt.Println("Exiting the program...")
}
