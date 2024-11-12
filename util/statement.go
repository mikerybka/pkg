package util

// Statement represents a line of code in a function or method body.
// There are 4 types of statements: returns, assignments, ifs and loops.
// Ifs and loops have substatements.
type Statement struct {
	IsReturn bool        `json:"isReturn"`
	Return   *Expression `json:"return"`

	IsAssign bool        `json:"isAssign"`
	Name     string      `json:"name"`
	Value    *Expression `json:"value"`

	IsIf      bool         `json:"isIf"`
	Condition *Expression  `json:"condition"`
	Body      []*Statement `json:"body"`

	// TODO: loops
}

func (st *Statement) Imports() ImportMap {
	imports := ImportMap{}

	if st.IsReturn {
		for k, v := range st.Return.Imports() {
			imports[k] = v
		}
	}

	if st.IsAssign {
		for k, v := range st.Value.Imports() {
			imports[k] = v
		}
	}

	if st.IsIf {
		for k, v := range st.Condition.Imports() {
			imports[k] = v
		}
	}

	return imports
}
