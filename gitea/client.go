package gitea

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mikerybka/pkg/git"
)

type Client struct {
	URL      string
	Username string
	Password string
}

func (c *Client) CloneOrPull(org, repo, dir string) error {
	// Use pkg/git
	gitClient := &git.Client{
		Dir: dir,
	}
	fi, err := os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		// If dir doesn't exist, clone
		return gitClient.Clone(c.URL)
	}
	if err != nil {
		// If there's some other read error, return it
		return err
	}
	if !fi.IsDir() {
		// If the dir is a file, report an error
		return fmt.Errorf("%s is a file, expected dir", dir)
	}
	// If the dir does exist, pull
	return gitClient.Pull()
}

func (c *Client) CreateRepo(org, repo string) error {
	url := fmt.Sprintf("%s/api/v1/orgs/%s/repos", c.URL, org)
	json := fmt.Sprintf("{\"name\":\"%s\",\"private\":true}", repo)
	body := bytes.NewReader([]byte(json))
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 201 {
		b, _ := io.ReadAll(res.Body)
		if len(b) > 0 {
			return fmt.Errorf("%s: %s", res.Status, b)
		} else {
			return fmt.Errorf(res.Status)
		}
	}
	return nil
}

// func (c *Client) WriteFile(org, repo, branch, path string, content []byte) error
