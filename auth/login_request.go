package auth

type LoginRequest struct {
	UserID string `json:"userID"`
	Code   string `json:"code"`
}
