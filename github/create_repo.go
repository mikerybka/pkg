package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func CreateRepo(org, repo, desc string, private bool) error {
	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos", org)

	// Encode the repo struct to JSON
	reqBody, err := json.Marshal(Repo{
		Name:        repo,
		Description: desc,
		Private:     private,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set the necessary headers
	req.Header.Set("Authorization", fmt.Sprintf("token %s", util.RequireEnvVar("GITHUB_TOKEN"))) // Authorization with Personal Access Token
	req.Header.Set("Content-Type", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-2XX status codes
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create repository, status code: %d", resp.StatusCode)
	}

	fmt.Println("Repository created successfully!")
	return nil
}
