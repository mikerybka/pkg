package hetzner

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func DeleteServer(id int) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("https://api.hetzner.cloud/v1/servers/%d", id), nil)
	req.Header.Add("Authorization", "Bearer "+util.RequireEnvVar("HETZNER_API_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf("%s: %s", res.Status, string(b))
	}

	return nil
}
