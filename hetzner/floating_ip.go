package hetzner

type FloatingIP struct {
	ID     int    `json:"id"`
	IP     string `json:"ip"`
	Server int    `json:"server"`
}
