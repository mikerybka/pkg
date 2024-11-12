package dev

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mikerybka/pkg/types"
)

func GetType(id string) (*types.Type, error) {
	url := fmt.Sprintf("https://data/types/raw/branch/main/%s.json", id)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http: %s", err)
	}
	var t types.Type
	err = json.NewDecoder(res.Body).Decode(&t)
	if err != nil {
		return nil, fmt.Errorf("json: %s", err)
	}
	return &t, nil
}
