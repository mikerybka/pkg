package brass

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
)

type Data struct {
	Dir string
}

func (d *Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check to see if a file exists at the requested path.
	path := filepath.Join(d.Dir, r.URL.Path)
	fi, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		panic(err)
	}

	// If the path is a directory, serve that.
	if fi.IsDir() {
		d := &Dir{
			Path: path,
		}
		d.ServeHTTP(w, r)
		return
	}

	// Otherwise, serve the file.
	f := &File{
		Path: path,
	}
	f.ServeHTTP(w, r)
}
