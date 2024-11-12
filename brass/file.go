package brass

import "net/http"

type File struct {
	Path string
}

func (f *File) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
