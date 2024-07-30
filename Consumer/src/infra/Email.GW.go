package infra

import (
	"log"
	"net/smtp"
)

type EmailSender struct {
	SMTPServer string
	Port       string
	Username   string
	Password   string
}

func NewEmailSender() *EmailSender {
	return &EmailSender{
		SMTPServer: "",
		Port:       "",
		Username:   "",
		Password:   "",
	}
}

func (e *EmailSender) SendEmail(to []string, subject, body string) error {
	from := e.Username
	auth := smtp.PlainAuth("", e.Username, e.Password, e.SMTPServer)
	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(e.SMTPServer+":"+e.Port, auth, from, to, msg)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Println("Email sent successfully to", to)
	return nil
}
