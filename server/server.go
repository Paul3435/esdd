package server

//Useful documentation: https://gobyexample.com/http-server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Paul3435/esdd/pkg/email"
)

type MailContent struct {
	Email   string
	Subject string
	Body    string
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, "landingPage.html")
}

func handleSendEmail(w http.ResponseWriter, r *http.Request) {
	//Parse the form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		return
	}

	//Render HTML
	renderHTML(w, "sendForm.html")

	//Gather data from the form in the POST req
	mail := &MailContent{
		Email:   r.PostFormValue("to"),
		Subject: r.PostFormValue("subject"),
		Body:    r.PostFormValue("body"),
	}

	emailServiceManager := email.NewEmailServiceManager(email.NewSendGrid("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"), email.NewMailgunService("02057f3a67aec395a2efd2e70426f144-911539ec-96815074"))
	emailServiceManager.SendEmail(mail.Subject, mail.Email, mail.Body)
}

func Start() {
	http.HandleFunc("/", landingPageHandler)
	http.HandleFunc("/send", handleSendEmail)
	http.ListenAndServe(":8080", nil)
}

func renderHTML(w http.ResponseWriter, htmlLocation string) error {
	html, err := os.ReadFile(htmlLocation)
	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, string(html))
	return nil
}
