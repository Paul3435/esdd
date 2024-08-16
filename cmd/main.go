package main

import (
	"fmt"
	"os"

	"github.com/Paul3435/esdd/pkg/tests"
	"github.com/Paul3435/esdd/server"
)

func main() {
	var mode string

	// Prompt the user to enter the mode
	fmt.Println("Please specify the mode: 'test' to run in testing mode or 'start' to run the server")
	fmt.Print("Enter mode: ")
	fmt.Scanln(&mode)

	// Switch based on the user input
	switch mode {
	case "test":
		tests.IntializeTests()
	case "start":
		server.Start()
	default:
		fmt.Println("Invalid mode. Please use 'test' or 'start'.")
		os.Exit(1)
	}
}
