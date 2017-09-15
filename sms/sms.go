package sms

import (
	"errors"
	"fmt"

	"github.com/sfreiberg/gotwilio"

	log "github.com/sirupsen/logrus"
)

const accountSid = "fill this in"
const authToken = "fill this in"
const from = "+1xxxyyyzzzz" // use twilio phone # here

var L = log.WithField("ctx", "sms")

var twilio *gotwilio.Twilio

func init() {
	twilio = gotwilio.NewTwilioClient(accountSid, authToken)
}

func Send(to, message string) error {
	resp, exception, err := twilio.SendSMS(from, to, message, "", "")
	L.Debug(fmt.Sprintf("resp: %+v", resp))
	L.Debug(fmt.Sprintf("exception: %+v", exception))

	if err != nil {
		return err
	}

	if exception != nil {
		err = errors.New(fmt.Sprintf("Status: %v, Message: %v", exception.Status, exception.Message))
	}

	return err
}
