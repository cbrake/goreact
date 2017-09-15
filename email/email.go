package email

import (
	"net/smtp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var L = log.WithField("ctx", "email")

func Send(to []string, subject, message string) error {
	// Set up authentication information.
	server := "smtp.mailgun.org"

	// fill in the user/pass with your auth data
	user := "postmaster@mg.mycompany.com"
	pass := "12341234123412341234"
	sender := "portal@mycompany.com"

	auth := smtp.PlainAuth("", user, pass, server)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := []byte("To: " + strings.Join(to, ", ") + "\r\n" +
		"Subject:" + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")

	err := smtp.SendMail(server+":587", auth, sender, to, msg)

	L.WithFields(log.Fields{
		"to":      to,
		"subject": subject,
		"message": message,
	}).Debug("sending email")

	return err
}
