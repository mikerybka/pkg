package commander

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Request struct {
	Cmd   string            `json:"cmd"`
	Dir   string            `json:"dir"`
	Env   map[string]string `json:"env"`
	Stdin string            `json:"stdin"`
}

func (req *Request) Handle() (*Response, error) {
	c := strings.Split(req.Cmd, " ")
	cmd := exec.Command(c[0], c[1:]...)

	if req.Dir != "" {
		cmd.Dir = req.Dir
	}

	for k, v := range req.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	cmd.Stdin = bytes.NewBufferString(req.Stdin)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	stdout, err := io.ReadAll(stdoutPipe)
	if err != nil {
		return nil, err
	}

	stderr, err := io.ReadAll(stderrPipe)
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	res := &Response{
		ExitCode: cmd.ProcessState.ExitCode(),
		Stdout:   string(stdout),
		Stderr:   string(stderr),
	}

	return res, nil
}
