package schemacafe

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

type Client struct {
}

func (c *Client) GetType(id string) (*util.Type, error) {
	url := fmt.Sprintf("https://api.schema.cafe/%s", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	t := &util.Type{}
	err = json.NewDecoder(resp.Body).Decode(t)
	return t, err
}
