package brass

import "net/http"

type Dir struct {
	Path string
}

func (d *Dir) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
