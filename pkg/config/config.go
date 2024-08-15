package config

import "os"

type Config struct {
	SendGridAPIKey string
	MailgunAPIKey  string
}

func LoadEnvVariables() *Config {
	return &Config{
		SendGridAPIKey: os.Getenv("SENDGRID_API_KEY"),
		MailgunAPIKey:  os.Getenv("MAILGUN_API_KEY"),
	}
}
