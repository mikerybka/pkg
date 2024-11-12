package hetzner

type CreateVMResponse struct {
	Server       *Server `json:"server"`
	RootPassword string  `json:"root_password"`
}
