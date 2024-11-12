package auth

type LogoutRequest struct {
	Session *Session `json:"session"`
}
