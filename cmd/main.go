package main

import (
	"os"

	"github.com/Paul3435/esdd/server"
)

func main() {
	os.Setenv("SENDGRID_API_KEY", "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")
	os.Setenv("MAILGUN_API_KEY", "02057f3a67aec395a2efd2e70426f144-911539ec-96815074")

	server.Start()
}
