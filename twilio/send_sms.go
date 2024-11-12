package twilio

import "github.com/mikerybka/pkg/util"

func SendSMS(to, message string) error {
	c := &Client{
		AccountSID:  util.RequireEnvVar("TWILIO_ACCOUNT_SID"),
		AuthToken:   util.RequireEnvVar("TWILIO_AUTH_TOKEN"),
		PhoneNumber: util.RequireEnvVar("TWILIO_PHONE_NUMBER"),
	}
	return c.SendSMS(to, message)
}
