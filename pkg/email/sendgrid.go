package email

//Official docu https://www.twilio.com/docs/sendgrid/for-developers/sending-email/quickstart-go#send-an-email

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	APIKey string
}

func NewSendGrid(apiKey string) *SendGrid {
	return &SendGrid{APIKey: apiKey}
}

func (s *SendGrid) SendEmail(subject, to, body string) {
	from := mail.NewEmail("Paul Borjesson", "Paul.borjesson.sesma@gmail.com")
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	toEmail := mail.NewEmail("Paul 2", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, htmlContent)
	client := sendgrid.NewSendClient(s.APIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
	} else {
		fmt.Println("Response: ", response.StatusCode)
	}

}
