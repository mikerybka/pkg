package gitea

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	SHA  string `json:"sha"`
}
