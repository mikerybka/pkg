package hetzner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func AssignFloatingIP(ip string, serverID int) error {
	allFloatingIPs, err := ListFloatingIPs()
	if err != nil {
		return err
	}

	for _, floatingIP := range allFloatingIPs {
		if floatingIP.IP == ip {
			b, _ := json.Marshal(AssignFloatingIPRequest{
				Server: serverID,
			})
			url := fmt.Sprintf("https://api.hetzner.cloud/v1/floating_ips/%d/actions/assign", floatingIP.ID)
			req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
			req.Header.Add("Authorization", "Bearer "+util.RequireEnvVar("HETZNER_API_TOKEN"))
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}

			if res.StatusCode != http.StatusCreated {
				b, _ := io.ReadAll(res.Body)
				return fmt.Errorf("%s: %s", res.Status, string(b))
			}

			return nil
		}
	}

	return fmt.Errorf("floating ip %s does not exist", ip)
}
