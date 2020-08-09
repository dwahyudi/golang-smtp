package util

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

/*
Mail is a generic struct type for representing a mail send request.
*/
type Mail struct {
	From       string
	To         []string
	Subject    string
	BodyType   string
	Body       string
	Attachment []string
}

/*
MailSend sends email with settings configured by envs.
*/
func MailSend(mail Mail) {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.From)
	m.SetHeader("To", mail.To...)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody(mail.BodyType, mail.Body)
	if len(mail.Attachment) > 0 {
		for _, attachment := range mail.Attachment {
			m.Attach(attachment)
		}
	}

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	CheckErr(err)
	d := gomail.NewDialer(os.Getenv("SMTP_HOSTNAME"),
		port,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"))

	err = d.DialAndSend(m)
	CheckErr(err)
}
