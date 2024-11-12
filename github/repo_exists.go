package github

import (
	"fmt"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

func RepoExists(org, repo string) (bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", org, repo)

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "token "+util.RequireEnvVar("GITHUB_TOKEN"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	return res.StatusCode == http.StatusOK, nil
}
