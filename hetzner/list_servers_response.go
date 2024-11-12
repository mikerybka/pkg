package hetzner

type ListServersResponse struct {
	Servers []Server `json:"servers"`
}
