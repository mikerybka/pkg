package hetzner

type CreateVMRequest struct {
	Name       string `json:"name"`
	ServerType string `json:"server_type"`
	Image      string `json:"image"`
	Location   string `json:"location"`
	SSHKeys    []int  `json:"ssh_keys"`
}
