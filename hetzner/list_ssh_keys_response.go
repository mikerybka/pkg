package hetzner

type ListSSHKeysResponse struct {
	SSHKeys []SSHKey `json:"ssh_keys"`
}
