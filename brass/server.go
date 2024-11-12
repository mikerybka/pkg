package brass

import (
	"net/http"
)

type Server struct {
	Data *Data
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: auth
	s.Data.ServeHTTP(w, r)
}
