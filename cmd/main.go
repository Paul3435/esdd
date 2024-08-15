package main

import "github.com/Paul3435/esdd/pkg/email"

func main() {
	sendGridAPIKey := "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"
	sendGrid := email.NewSendGrid(sendGridAPIKey)

	sendGrid.SendEmail("Test Subject", "paul.borjesson.sesma3435@gmail.com", "This is a test email body.")

	mailGunAPIKey := "02057f3a67aec395a2efd2e70426f144-911539ec-96815074"
	mailGun := email.NewMailgunService(mailGunAPIKey)

	mailGun.SendEmail("Test Subject", "paul.borjesson.sesma3435@gmail.com", "This is a test email body.")
}
