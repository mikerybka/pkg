package auth

import "github.com/mikerybka/pkg/util"

type SendLoginCodeRequest struct {
	Phone util.PhoneNumber `json:"phone"`
}
