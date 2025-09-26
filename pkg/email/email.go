package emails

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmailsTo(to string, subject string, body string) error {
	log.Println("Sending email to:", to)
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	mail := gomail.NewMessage()
	mail.SetHeader("From", os.Getenv("SMTP_USER"))
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", body)
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(mail); err != nil {
		log.Println("Failed to send email", err.Error())
		return err
	}
	return nil
}
