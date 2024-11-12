package util

import (
	_ "embed"
	"strings"
)

type Function struct {
	ID       string      `json:"id"`
	Desc     string      `json:"desc"`
	Comments []Comment   `json:"comments"`
	Name     Name        `json:"name"`
	Inputs   []Field     `json:"inputs"`
	Outputs  []Field     `json:"outputs"`
	Body     []Statement `json:"body"`
}

func (f *Function) Imports() ImportMap {
	imports := ImportMap{}

	for _, in := range f.Inputs {
		typ := in.Type
		for {
			if strings.HasPrefix(typ, "[]") {
				typ = strings.TrimPrefix(typ, "[]")
			} else if strings.HasPrefix(typ, "map[string]") {
				typ = strings.TrimPrefix(typ, "map[string]")
			} else {
				break
			}
		}
		from, _ := parseName(typ)
		if from != "" {
			imports[from] = importPath(from)
		}
	}

	for _, out := range f.Outputs {
		typ := out.Type
		for {
			if strings.HasPrefix(typ, "[]") {
				typ = strings.TrimPrefix(typ, "[]")
			} else if strings.HasPrefix(typ, "map[string]") {
				typ = strings.TrimPrefix(typ, "map[string]")
			} else {
				break
			}
		}
		from, _ := parseName(typ)
		if from != "" {
			imports[from] = importPath(from)
		}
	}

	for _, st := range f.Body {
		for k, v := range st.Imports() {
			imports[k] = v
		}
	}

	return imports
}
