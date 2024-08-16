package tests

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Paul3435/esdd/pkg/email"
)

type EmailContent struct {
	Email          string
	Subject        string
	Body           string
	APIKeySendGrid string
	APIKeyMailgun  string
}

// func NewTestOnlyMail(email string, subject string, body string) *EmailContent {
// 	cfg := config.LoadEnvVariables()
// 	return &EmailContent{
// 		Email:          email,
// 		Subject:        subject,
// 		Body:           body,
// 		APIKeySendGrid: cfg.SendGridAPIKey,
// 		APIKeyMailgun:  cfg.MailgunAPIKey,
// 	}
// }

func IntializeTests() {
	choice := -1
	var emailContent = &EmailContent{} //empty Email Content, avoids null pointers should the user not fill the fields

	for choice != 0 {
		fmt.Println("\nWelcome to the command-line based testing environment.\n")
		fmt.Println("Select an option:")
		fmt.Println("1. Create email object (Email, Subject, Body), which will autopopulate the API Keys")
		fmt.Println("2. Create email object (Email, Subject, Body, ApiKeySendGrid, ApiKeyMailgun)")
		fmt.Println("3. Send email using only SendGrid")
		fmt.Println("4. Send email using only Mailgun")
		fmt.Println("5. Send email using both")
		fmt.Println("0. Exit")

		//Read input from user
		choice = selectOption()
		emailContent.triggerOption(choice)
	}

}

func (e *EmailContent) triggerOption(choice int) {
	switch choice {
	case 0:
		fmt.Println("Closing...")
	case 1:
		*e = *createEmailContent(false)
		fmt.Println("Created email object with autopopulated API Keys:")
		fmt.Printf("%+v\n", e)
	case 2:
		*e = *createEmailContent(true)
		fmt.Println("Created email object with custom API Keys:")
		fmt.Printf("%+v\n", e)
	case 3:
		fmt.Printf("e: %v\n", e)
		e.sendUsingSendGrid() //You can try sending emails with an empty emailContent
	case 4:
		e.sendUsingMailgun()
	case 5:
		e.sendUsingBoth()
	default:
		fmt.Println("Invalid choice.")
	}
}

// sendUsingSendGrid handles sending the email using the SendGrid service
func (e *EmailContent) sendUsingSendGrid() {
	fmt.Println(e)
	sendEmailWithService(
		email.NewSendGrid(e.APIKeySendGrid),
		e.Subject,
		e.Email,
		e.Body,
	)
}

// sendUsingMailgun handles sending the email using the Mailgun service
func (e *EmailContent) sendUsingMailgun() {
	sendEmailWithService(
		email.NewMailgunService(e.APIKeyMailgun),
		e.Subject,
		e.Email,
		e.Body,
	)
}

// sendUsingBoth handles sending the email using both SendGrid and Mailgun
func (e *EmailContent) sendUsingBoth() {
	emailServiceManager := email.NewEmailServiceManager(
		email.NewSendGrid(e.APIKeySendGrid),
		email.NewMailgunService(e.APIKeyMailgun),
	)

	emailServiceManager.SendEmail(e.Subject, e.Email, e.Body)
}

func sendEmailWithService(service email.EmailService, subject, to, body string) {
	err := service.SendEmail(subject, to, body)
	if err != nil {
		fmt.Printf("Failed to send email via %T: %v\n", service, err)
	} else {
		fmt.Printf("Email sent successfully via %T\n", service)
	}
}

// prompts the user to enter email details and returns an EmailContent struct
func createEmailContent(includeApiKeys bool) *EmailContent {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter Subject: ")
	subject, _ := reader.ReadString('\n')
	subject = strings.TrimSpace(subject)

	fmt.Print("Enter Body: ")
	body, _ := reader.ReadString('\n')
	body = strings.TrimSpace(body)

	emailContent := &EmailContent{
		Email:   email,
		Subject: subject,
		Body:    body,
	}

	if includeApiKeys {
		fmt.Print("Enter SendGrid API Key: ")
		apiKeySendGrid, _ := reader.ReadString('\n')
		apiKeySendGrid = strings.TrimSpace(apiKeySendGrid)

		fmt.Print("Enter Mailgun API Key: ")
		apiKeyMailgun, _ := reader.ReadString('\n')
		apiKeyMailgun = strings.TrimSpace(apiKeyMailgun)

		emailContent.APIKeySendGrid = apiKeySendGrid
		emailContent.APIKeyMailgun = apiKeyMailgun
	} else {
		emailContent.APIKeySendGrid = os.Getenv("SENDGRID_API_KEY")
		emailContent.APIKeyMailgun = os.Getenv("MAILGUN_API_KEY")
	}

	return emailContent
}

func selectOption() int {
	var choice int
	for {
		var input string

		fmt.Scanln(&input)

		var err error
		choice, err = strconv.Atoi(input)
		if err != nil || choice < 0 || choice > 5 {
			fmt.Println("Invalid choice. Please try again.")
		} else {
			break
		}
	}
	return choice
}
