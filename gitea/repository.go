package gitea

import "strings"

type Repository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	CloneURL string `json:"clone_url"`
}

func (r *Repository) Path() string {
	return strings.TrimSuffix(strings.TrimPrefix(r.CloneURL, "https://"), ".git")
}
