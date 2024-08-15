package server

//Useful documentation: https://gobyexample.com/http-server

import (
	"fmt"
	"net/http"

	"github.com/Paul3435/esdd/pkg/email"
)

type Server struct {
	emailManager *email.EmailServiceManager
}

func NewServer(emailManager *email.EmailServiceManager) *Server {
	return &Server{emailManager: emailManager}
}

func (s *Server) handleSendEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func (s *Server) Start() {
	http.HandleFunc("/send", s.handleSendEmail)
	http.ListenAndServe(":8080", nil)
}
