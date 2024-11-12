package util

type Field struct {
	ID       string    `json:"id"`
	Desc     string    `json:"desc"`
	Comments []Comment `json:"comments"`
	Name     Name      `json:"name"`
	Type     string    `json:"type"`
}
