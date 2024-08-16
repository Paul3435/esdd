package main

import (
	"fmt"
	"strings"

	"github.com/Paul3435/esdd/pkg/tests"
	"github.com/Paul3435/esdd/server"
)

func main() {
	var input string
	fmt.Print("Enter 'test' to run in testing mode or 'start' to run the server: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)
	if strings.ToLower(input) == "test" {
		tests.IntializeTests()
	} else if strings.ToLower(input) == "start" {
		server.Start()
	} else {
		fmt.Println("Invalid input. Please enter 'test' or 'start'.")
	}
}
