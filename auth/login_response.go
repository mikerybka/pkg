package auth

type LoginResponse struct {
	Ok     bool   `json:"ok"`
	Error  string `json:"error"`
	UserID string `json:"userID"`
	Token  string `json:"token"`
}
