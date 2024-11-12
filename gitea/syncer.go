package gitea

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/mikerybka/pkg/git"
	"github.com/mikerybka/pkg/util"
)

type Syncer struct {
	SrcDir string
}

func (s *Syncer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	push := &Push{}
	err := json.NewDecoder(r.Body).Decode(push)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !strings.HasPrefix(push.Ref, "refs/heads/") {
		return
	}
	git := &git.Client{
		Dir: filepath.Join(s.SrcDir, push.Branch(), push.Repository.Path()),
	}
	if !util.IsDir(git.Dir) {
		err := git.Clone(push.Repository.CloneURL)
		if err != nil {
			panic(err)
		}
		err = git.Checkout(push.Branch())
		if err != nil {
			panic(err)
		}
	} else {
		err := git.Pull()
		if err != nil {
			panic(err)
		}
	}
}
