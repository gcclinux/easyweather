package main

import (
	"fmt"
	"os"
	"path/filepath"
	"weatherCollect/handlers"
)

func main() {

	var argument = len(os.Args)

	if argument == 1 {
		displayMenu()
	} else if argument == 2 && os.Args[1] == "--help" {
		displayMenu()
	} else if argument == 2 && os.Args[1] == "--collect" {
		handlers.CollectLoop()
	} else if argument == 2 && os.Args[1] == "--web" {
		handlers.LaunchWeb()
	} else if argument == 2 && os.Args[1] == "--integrity" {
		fmt.Println(handlers.SetupIntegraty())
	} else {
		displayMenu()
	}

}

func displayMenu() {

	ex := filepath.Base(os.Args[0])
	fmt.Println("--------------------------------")
	fmt.Println("Available Options")
	fmt.Println()
	fmt.Println("# Display the Help Menu")
	fmt.Print("$ ", ex, " --help\n\n")
	fmt.Println("# Launch the web service to open Weather Page!")
	fmt.Print("$ ", ex, " --web\n\n")
	fmt.Println("# Start collecting weather and put it in background!")
	fmt.Print("$ ", ex, " --collect\n\n")
	fmt.Println("# Run a Setup integraty check before starting or in case of errors!")
	fmt.Print("$ ", ex, " --integrity\n\n")
}
