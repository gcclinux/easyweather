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
	fmt.Println("# Start Collecting weather in the backgound!")
	fmt.Print("$ ", ex, " --collect\n\n")
}
