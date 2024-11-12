package hetzner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func CreateVM(name string) (*Server, error) {
	sshKeys, err := ListSSHKeys()
	if err != nil {
		return nil, err
	}

	b, _ := json.Marshal(CreateVMRequest{
		Name:       name,
		ServerType: "ccx13",
		Image:      "ubuntu-24.04",
		Location:   "ash",
		SSHKeys:    sshKeys,
	})
	req, _ := http.NewRequest("POST", "https://api.hetzner.cloud/v1/servers", bytes.NewReader(b))
	req.Header.Add("Authorization", "Bearer "+util.RequireEnvVar("HETZNER_API_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("%s: %s", res.Status, string(b))
	}

	r := &CreateVMResponse{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r.Server, nil
}
