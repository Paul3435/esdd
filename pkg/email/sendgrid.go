package email

//Official docu https://www.twilio.com/docs/sendgrid/for-developers/sending-email/quickstart-go#send-an-email

import (
	"errors"

	"github.com/sendgrid/sendgrid-go"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	APIKey string
}

func NewSendGrid(apiKey string) *SendGrid {
	return &SendGrid{APIKey: apiKey}
}

func (s *SendGrid) SendEmail(subject, to, body string) error {
	//Contents of mail
	from := mail.NewEmail("Paul Borjesson", "Paul.borjesson.sesma@gmail.com")
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	toEmail := mail.NewEmail("Paul 2", to)
	message := mail.NewSingleEmail(from, subject, toEmail, body, htmlContent)

	//Client
	client := sendgrid.NewSendClient(s.APIKey)
	//Response
	response, err := client.Send(message)

	//Handling response. This library usually doesn't properly log errors, so using isCodeValid() is needed.
	if err != nil || !isCodeValid(response.StatusCode) {
		return errors.New("Failed to send email: " + response.Body)
	}

	return nil
}

// Verifies the most significative value by dividing the number iteratively
func isCodeValid(responseCode int) bool {
	var i int
	for i = responseCode; i >= 10; i = i / 10 {
	}
	return i == 2
}

func (s *SendGrid) Name() string {
	return "SendGrid"
}
