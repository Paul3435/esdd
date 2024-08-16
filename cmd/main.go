package main

import (
	"os"

	"github.com/Paul3435/esdd/pkg/tests"
	"github.com/Paul3435/esdd/server"
)

func main() {
	//os.Setenv("TEST_MODE", "true")
	testing := os.Getenv("TEST_MODE")

	if testing == "true" {
		tests.IntializeTests()
	} else {
		server.Start()
	}
}
