package main

import (
	"github.com/Paul3435/esdd/pkg/email"
)

func main() {
	// sendGridAPIKey := "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"
	// sendGrid := email.NewSendGrid(sendGridAPIKey)

	// err1 := sendGrid.SendEmail("Test Subject", "paul.borjesson.sesma3435@gmail.com", "This is a test email body.")
	// fmt.Println("ERROR FOUND:", err1)

	// mailGunAPIKey := "02057f3a67aec395a2efd2e70426f144-911539ec-96815074" //Later on they will be substracted from the os
	// mailGun := email.NewMailgunService(mailGunAPIKey)

	// err2 := mailGun.SendEmail("Test Subject", "paul.borjesson.sesma3435@gmail.com", "This is a test email body.")
	// fmt.Println("ERROR FOUND:", err2)

	emailServiceManager := email.NewEmailServiceManager(email.NewSendGrid("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"), email.NewMailgunService("02057f3a67aec395a2efd2e70426f144-911539ec-96815074"))
	emailServiceManager.SendEmail("Test Subject", "paul.borjesson.sesma3435@gmail.com", "This is a test email body.")
}
