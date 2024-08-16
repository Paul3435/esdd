package main

import (
	"fmt"

	"github.com/Paul3435/esdd/pkg/tests"
	"github.com/Paul3435/esdd/server"
)

func main() {
	var mode string

	// Prompt the user to enter the mode
	fmt.Println("Write 'test' if you wish to enter testing mode. Else skip.")
	fmt.Print("Enter mode: ")
	fmt.Scanln(&mode)

	// Switch based on the user input
	switch mode {
	case "test":
		tests.IntializeTests()
	default:
		server.Start()
	}
}
