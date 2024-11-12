package util

type Package struct {
	ID        string              `json:"id"`
	Desc      string              `json:"desc"`
	Comments  []Comment           `json:"comments"`
	Types     map[string]Type     `json:"types"`
	Functions map[string]Function `json:"functions"`
	Variables map[string]Value    `json:"variables"`
	Constants map[string]Value    `json:"constants"`
}
