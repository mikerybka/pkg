package gitea

import "strings"

type Push struct {
	Ref        string      `json:"ref"`
	Before     string      `json:"before"`
	After      string      `json:"after"`
	CompareURL string      `json:"compare_url"`
	Repository *Repository `json:"repository"`
}

func (p *Push) Branch() string {
	return strings.TrimPrefix(p.Ref, "refs/heads/")
}
