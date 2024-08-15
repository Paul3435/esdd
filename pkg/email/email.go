package email

//This interface will be implemented by all the email delivering systems, this allows for both resilience and easily cycling through them for fallbacks
type EmailService interface {
	SendEmail(subject, to, body string) error
}
