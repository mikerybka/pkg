package dev

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mikerybka/pkg/gitea"
)

func ListTypeIDs() ([]string, error) {
	req, err := http.NewRequest("GET", "https://api/v1/repos/data/types/contents", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var files []gitea.File
	err = json.NewDecoder(res.Body).Decode(&files)
	if err != nil {
		return nil, err
	}
	typeIDs := []string{}
	for _, f := range files {
		id := strings.TrimSuffix(f.Name, ".json")
		typeIDs = append(typeIDs, id)
	}
	return typeIDs, nil
}
