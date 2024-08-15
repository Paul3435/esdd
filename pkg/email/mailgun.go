package email

//Official docu: https://documentation.mailgun.com/docs/mailgun/sdk/go_sdk/
//Relevant: https://pkg.go.dev/github.com/mailgun/mailgun-go/v4#section-readme

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunService struct {
	Domain string
	APIKey string
}

func NewMailgunService(apiKey string) *MailgunService {
	return &MailgunService{
		Domain: "sandbox404e5becda2f49408e32649074e624e8.mailgun.org",
		APIKey: apiKey,
		//Url:    "https://api.eu.mailgun.net/v3", //My account is based in the US, so not necessary
	}
}

func (m *MailgunService) SendEmail(subject, to, body string) {
	mg := mailgun.NewMailgun(m.Domain, m.APIKey)
	//mg.SetAPIBase(m.Url)

	senderEmail := "Paul <mailgun@" + m.Domain + ">"
	fmt.Println(senderEmail)

	message := mg.NewMessage(
		senderEmail,
		subject,
		body,
		to,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, id, err := mg.Send(ctx, message)
	if err != nil {
		fmt.Println("Response: ", response, "Id: ", id)
		log.Fatalf("Failed to send email: %v", err)
	} else {
		fmt.Println("Success: ", response)
	}
}
