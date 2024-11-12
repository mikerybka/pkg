package util

import "time"

type Comment struct {
	Timestamp time.Time `json:"timestamp"`
	Author    string    `json:"author"`
	Message   string    `json:"message"`
	Replies   []Comment `json:"replies"`
}
