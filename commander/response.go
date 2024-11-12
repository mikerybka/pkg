package commander

type Response struct {
	ExitCode int    `json:"exitCode"`
	Stderr   string `json:"stderr"`
	Stdout   string `json:"stdout"`
}
