package hetzner

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func ListSSHKeys() ([]int, error) {
	req, _ := http.NewRequest("GET", "https://api.hetzner.cloud/v1/ssh_keys", nil)
	req.Header.Add("Authorization", "Bearer "+util.RequireEnvVar("HETZNER_API_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("%s: %s", res.Status, string(b))
	}

	r := &ListSSHKeysResponse{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	ids := []int{}
	for _, key := range r.SSHKeys {
		ids = append(ids, key.ID)
	}

	return ids, nil
}
