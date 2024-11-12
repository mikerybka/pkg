package hetzner

type Server struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PublicNet struct {
		IPv4 struct {
			IP string `json:"ip"`
		} `json:"ipv4"`
	} `json:"public_net"`
}
