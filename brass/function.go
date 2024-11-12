package brass

import (
	_ "embed"
	"net/http"

	"github.com/mikerybka/pkg/util"
)

type Function struct {
	Inputs  []Field     `json:"inputs"`
	Outputs []Field     `json:"outputs"`
	Body    []Statement `json:"body"`
}

func (fn *Function) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		util.WriteJSON(w, Value{
			Type: "Function",
			Data: fn,
		})
		return
	}

	if r.Method == "POST" {
		// Parse
		req := util.ReadJSON[util.MethodCall](r.Body)

		// Execute
		res, err := util.CallMethod(fn, req.Method, req.Args)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond
		util.WriteJSON(w, res)
		return
	}
}
