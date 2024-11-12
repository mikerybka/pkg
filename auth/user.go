package auth

import (
	"fmt"

	"github.com/mikerybka/pkg/util"
)

type User struct {
	ID string

	Phone util.PhoneNumber

	LoginCode     string
	SessionTokens util.Set[string]
}

func (u *User) Authenticated() bool {
	return u.ID != ""
}

func (u *User) ValidSession(token string) bool {
	return u.Authenticated() && u.SessionTokens.Has(token)
}

func (u *User) EndSession(token string) {
	u.SessionTokens.Remove(token)
}

func (u *User) SendLoginCode(twilio *util.TwilioClient) error {
	code := util.RandomCode(6)
	msg := fmt.Sprintf("Your login code is %s", code)
	err := twilio.SendSMS(u.Phone, msg)
	if err != nil {
		return err
	}
	u.LoginCode = code
	return nil
}

func (u *User) Login(code string) (string, error) {
	if code != u.LoginCode {
		return "", fmt.Errorf("wrong login code")
	}

	token := util.RandomToken(32)
	u.SessionTokens.Add(token)
	return token, nil
}

func (u *User) Logout(token string) {
	u.SessionTokens.Remove(token)
}
