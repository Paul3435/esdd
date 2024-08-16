package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Paul3435/esdd/pkg/tests"
	"github.com/Paul3435/esdd/server"
)

func main() {
	mode := flag.String("mode", "start", "Specify the mode: 'test' to run in testing mode or 'start' to run the server")
	flag.Parse()

	switch *mode {
	case "test":
		tests.IntializeTests()
	case "start":
		server.Start()
	default:
		fmt.Println("Invalid mode. Please use 'test' or 'start'.")
		os.Exit(1)
	}
}
