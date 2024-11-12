package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Client struct {
	Dir string
}

func (c *Client) HasChanges() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = c.Dir
	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("failed to run git status: %v", err)
	}

	// Trim the output and check if it's empty
	changes := strings.TrimSpace(string(out))

	if changes == "" {
		return false, nil // No changes
	}
	return true, nil // There are changes
}

func (c *Client) Clone(url string) error {
	dir := filepath.Dir(c.Dir)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	cmd := exec.Command("git", "clone", url)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}
	return nil
}

func (c *Client) Checkout(branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = c.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}
	return nil
}

func (c *Client) AddAll() error {
	cmd := exec.Command("git", "add", "--all")
	cmd.Dir = c.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}
	return nil
}

func (c *Client) Pull() error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = c.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}
	return nil
}

func (c *Client) Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Dir = c.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}
	return nil
}

func (c *Client) Push() error {
	cmd := exec.Command("git", "push")
	cmd.Dir = c.Dir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, out)
	}
	return nil
}
