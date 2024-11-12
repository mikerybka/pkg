package backend

import (
	"github.com/mikerybka/pkg/twilio"
	"github.com/mikerybka/pkg/util"
)

var adminPhone = util.RequireEnvVar("ADMIN_PHONE")

func Notify(msg string) error {
	return twilio.SendSMS(adminPhone, msg)
}
