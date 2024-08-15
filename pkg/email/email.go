package email

import "fmt"

//This interface will be implemented by all the email delivering systems, this allows for both resilience and easily cycling through them for fallbacks
type EmailService interface {
	SendEmail(subject, to, body string) error
	Name() string //Each provider is granted a name for comfortability in logging
}

//Holds the collection of providers.
type EmailServiceManager struct {
	providers []EmailService
}

//Instantiates EmailServiceManager, in the future you should be able to select which providers you want to use
func NewEmailServiceManager(providers ...EmailService) *EmailServiceManager {
	return &EmailServiceManager{providers: providers}
}

//If a provider in the range fails, tries with the next.
func (m *EmailServiceManager) SendEmail(to, subject, body string) {
	for i, provider := range m.providers {
		err := provider.SendEmail(to, subject, body)
		if err == nil {
			fmt.Println("Email sent with", m.providers[i].Name(), "successfully.")
			return
		}
		fmt.Println("Service:", m.providers[i].Name(), "failed with: \n", err)
	}
}
