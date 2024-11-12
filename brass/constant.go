package brass

import (
	_ "embed"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

type Constant struct {
	Type  *Type `json:"type"`
	Value any   `json:"value"`
}

func (c *Constant) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		util.WriteJSON(w, Value{
			Type: "Constant",
			Data: c,
		})
		return
	}

	if r.Method == "POST" {
		// Parse
		req := util.ReadJSON[util.MethodCall](r.Body)

		// Execute
		res, err := util.CallMethod(c, req.Method, req.Args)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond
		util.WriteJSON(w, res)
		return
	}
}
