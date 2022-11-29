package services

import (
	"fmt"
	"github.com/spf13/viper"
	"net/smtp"
	"os"
	"strings"
)

type Mail struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func SendMail(m Mail) error {
	smtpEmail := os.Getenv("SMTP_EMAIL")
	smtpPassword := os.Getenv("SMTP_EMAIL_PASSWORD")

	smtpHost := viper.GetString("smtp.host")
	smtpPort := viper.GetString("rabbitmq.port")
	smtpAddress := smtpHost + ":" + smtpPort

	// Building the message
	message := buildMessage(m)

	// Building the auth data
	auth := smtp.PlainAuth("", smtpEmail, smtpPassword, smtpHost)

	// Sending mail
	err := smtp.SendMail(smtpAddress, auth, smtpEmail, m.To, []byte(message))
	return err
}

func buildMessage(m Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	//msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(m.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", m.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", m.Body)
	return msg
}
