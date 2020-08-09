package emailing

import (
	"github.com/dwahyudi/golang-smtp/util"
)

/*
RegistrationWelcomeSend send registration email to designated email address.
*/
func RegistrationWelcomeSend(emailAddress string) {
	mail := util.Mail{
		From:       "no-reply@hogwartz.com",
		To:         []string{emailAddress},
		Subject:    "Welcome to Hogwartz",
		BodyType:   "text/html",
		Body:       registrationMailBody(),
		Attachment: []string{"temp/hogwartz.jpg", "temp/owl.jpg", "temp/apprentice-equip-list"},
	}

	util.MailSend(mail)
}

func registrationMailBody() string {
	return "<html><body><p>We pleased to inform you that you have a place at Hogwartz School of Witchcraft and Wizardry. <br/>Please find enclosed a list of necessary books and equipments.</p></body></html>"
}
