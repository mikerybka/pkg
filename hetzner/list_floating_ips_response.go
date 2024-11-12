package hetzner

type ListFloatingIPsResponse struct {
	FloatingIPs []FloatingIP `json:"floating_ips"`
}
