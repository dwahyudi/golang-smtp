package demo

import "github.com/dwahyudi/golang-smtp/util"

func SimpleMailDemo() {
	mail := util.Mail{
		From:       "no-reply@example00.com",
		To:         []string{"user01@example99.com", "user02@example.com"},
		Subject:    "Sample Subject",
		BodyType:   "text/html",
		Body:       "<html><body><p>Sample body</p></body></html>",
		Attachment: []string{"temp/cat.jpg", "temp/orange.jpg"},
	}

	util.MailSend(mail)
}
