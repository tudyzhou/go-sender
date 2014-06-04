package sender

import (
	"github.com/tudyzhou/go-sender/config"
	"net/smtp"
)

type EmailInfo struct {
	To       string
	Subject  string
	Body     string
	MailType string
}

// sender model
type sender struct {
	EmailLis []config.EmailModel
}

var (
	Sender = &sender{config.EmailConfig}
)

func (api *sender) Send(e *EmailInfo) (err error) {
	// auto pick one senderEmail
	s := api.EmailLis[0]

	// content_type
	var contentType string
	switch e.MailType {
	case "html":
		contentType = "Content-Type: text/html; charset=UTF-8"
	default:
		contentType = "Content-Type: text/plain; charset=UTF-8"
	}

	// combine msg
	msg := []byte("To: " + e.To + "\r\n" +
		"From: " + s.User + "<" + s.User + ">\r\n" +
		"Subject: " + e.Subject + "\r\n" +
		contentType + "\r\n\r\n" +
		e.Body)

	auth := smtp.PlainAuth(
		"",
		s.User,
		s.Password,
		s.Host,
	)
	err = smtp.SendMail(
		s.Host+":"+s.Port,
		auth,
		s.User,
		[]string{e.To},
		msg,
	)
	return
}
