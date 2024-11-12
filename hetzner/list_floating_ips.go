package hetzner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func ListFloatingIPs() ([]FloatingIP, error) {
	req, _ := http.NewRequest("GET", "https://api.hetzner.cloud/v1/floating_ips", nil)
	req.Header.Add("Authorization", "Bearer "+util.RequireEnvVar("HETZNER_API_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("%s: %s", res.Status, string(b))
	}

	r := &ListFloatingIPsResponse{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	return r.FloatingIPs, nil
}
