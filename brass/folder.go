package brass

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/mikerybka/pkg/util"
)

type Folder struct {
	Constants  map[string]*Constant `json:"constants"`
	Variables  map[string]*Variable `json:"variables"`
	Types      map[string]*Type     `json:"types"`
	Functions  map[string]*Function `json:"functions"`
	SubFolders map[string]*Folder   `json:"subFolders"`
}

// List returns a list of items in the folder sorted by ID.
func (f *Folder) List() []ListItem {
	items := []ListItem{}

	for id := range f.Constants {
		items = append(items, ListItem{
			ID:   id,
			Type: "Constant",
		})
	}

	for id := range f.Variables {
		items = append(items, ListItem{
			ID:   id,
			Type: "Variable",
		})
	}

	for id := range f.Types {
		items = append(items, ListItem{
			ID:   id,
			Type: "Type",
		})
	}

	for id := range f.Functions {
		items = append(items, ListItem{
			ID:   id,
			Type: "Function",
		})
	}

	for id := range f.SubFolders {
		items = append(items, ListItem{
			ID:   id,
			Type: "Folder",
		})
	}

	slices.SortFunc(items, func(a ListItem, b ListItem) int {
		if a.ID < b.ID {
			return -1
		}
		if a.ID > b.ID {
			return 1
		}
		// a.ID == b.ID
		return 0
	})

	return items
}

func (f *Folder) Exists(id string) bool {
	_, ok := f.Constants[id]
	if ok {
		return true
	}

	_, ok = f.Variables[id]
	if ok {
		return true
	}

	_, ok = f.Types[id]
	if ok {
		return true
	}

	_, ok = f.Functions[id]
	if ok {
		return true
	}

	_, ok = f.SubFolders[id]
	return ok
}

func (f *Folder) Add(kind, id, data string) error {
	if f.Exists(id) {
		return fmt.Errorf("%s aleady exists", id)
	}

	switch kind {
	case "Folder":
		d := &Folder{}
		err := json.Unmarshal([]byte(data), d)
		if err != nil {
			return err
		}
		f.SubFolders[id] = d
		return nil
	case "Constant":
		c := &Constant{}
		err := json.Unmarshal([]byte(data), c)
		if err != nil {
			return err
		}
		f.Constants[id] = c
		return nil
	case "Variable":
		v := &Variable{}
		err := json.Unmarshal([]byte(data), v)
		if err != nil {
			return err
		}
		f.Variables[id] = v
		return nil
	case "Type":
		t := &Type{}
		err := json.Unmarshal([]byte(data), t)
		if err != nil {
			return err
		}
		f.Types[id] = t
		return nil
	case "Function":
		fn := &Function{}
		err := json.Unmarshal([]byte(data), fn)
		if err != nil {
			return err
		}
		f.Functions[id] = fn
		return nil
	default:
		return fmt.Errorf("unkown kind: %s", kind)
	}
}

func (f *Folder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	first, rest, isRoot := util.PopPath(r.URL.Path)
	if isRoot {
		if r.Method == "GET" {
			util.WriteJSON(w, Value{
				Type: "brass.Folder",
				Data: f,
			})
			return
		}
		if r.Method == "POST" {
			req := util.ReadJSON[util.MethodCall](r.Body)
			switch req.Method {
			case "Add":
				err := f.Add(
					util.ParseJSON[string](req.Args[0]),
					util.ParseJSON[string](req.Args[1]),
					util.ParseJSON[string](req.Args[2]),
				)
				json.NewEncoder(w).Encode([]string{
					util.JSONString(err),
				})
				return
			default:
				http.Error(w, "unkown method "+req.Method, http.StatusBadRequest)
				return
			}
		}
	}

	if rest == "/" {
		c, ok := f.Constants[first]
		if ok {
			c.ServeHTTP(w, r)
			return
		}

		v, ok := f.Variables[first]
		if ok {
			v.ServeHTTP(w, r)
			return
		}

		t, ok := f.Types[first]
		if ok {
			t.ServeHTTP(w, r)
			return
		}

		fn, ok := f.Functions[first]
		if ok {
			fn.ServeHTTP(w, r)
			return
		}
	}

	f, ok := f.SubFolders[first]
	if ok {
		r.URL.Path = rest
		f.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}
