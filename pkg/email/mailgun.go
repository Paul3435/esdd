package email

//Official docu: https://documentation.mailgun.com/docs/mailgun/sdk/go_sdk/
//Relevant: https://pkg.go.dev/github.com/mailgun/mailgun-go/v4#section-readme

import (
	"context"

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

func (m *MailgunService) SendEmail(subject, to, body string) error {
	//Instantiate MailGun service
	mg := mailgun.NewMailgun(m.Domain, m.APIKey)
	//mg.SetAPIBase(m.Url)  //Necessary if the account is european

	//Contents of mail
	senderEmail := "Paul <mailgun@" + m.Domain + ">"
	message := mg.NewMessage(
		senderEmail,
		subject,
		body,
		to,
	)

	//Mailgun works by enqueuing messages, and requires context as parameter.
	ctx := context.Background()

	_, _, err := mg.Send(ctx, message)

	return err
}

func (m *MailgunService) Name() string {
	return "Mailgun"
}
