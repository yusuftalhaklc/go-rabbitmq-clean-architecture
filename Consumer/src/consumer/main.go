package main

import (
	"encoding/json"
	"infra"
	"log"
)

func main() {
	rmq := infra.NewRabbitMQConsumer()
	emailgw := infra.NewEmailSender()

	rmq.StartConsuming(func(body []byte) {
		var template EmailTemplate

		err := json.Unmarshal(body, &template)
		if err != nil {
			log.Println("Consume Body Error")
		}

		switch template.Type {
		case Verification:
			err := emailgw.SendEmail([]string{template.To}, "Email Verification", template.Link)
			if err != nil {
				log.Println("Verification Email send Error")
			}
		case PasswordReset:
			log.Println("Password Reset Send")
		default:
			log.Println("Do not know type")
		}

	})
}

type EmailTemplate struct {
	To   string `json:"to"`
	Link string `json:"link"`
	Type string `json:"type"`
}

const (
	Verification  = "verification"
	PasswordReset = "passwordReset"
)
