package server

//Useful documentation: https://gobyexample.com/http-server

import (
	"fmt"
	"net/http"
	"os"
)

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("./landingPage.html")
	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, string(html))
}

func handleSendEmail(w http.ResponseWriter, r *http.Request) {

}

func Start() {
	http.HandleFunc("/", landingPageHandler)
	http.HandleFunc("/send", handleSendEmail)
	http.ListenAndServe(":8080", nil)
}
