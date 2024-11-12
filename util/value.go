package util

type Value struct {
	ID       string    `json:"id"`
	Desc     string    `json:"desc"`
	Comments []Comment `json:"comments"`
	Type     string    `json:"type"`
	JSON     string    `json:"json"`
}
