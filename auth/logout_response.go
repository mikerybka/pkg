package auth

type LogoutResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}
