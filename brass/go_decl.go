package brass

type GoDecl struct {
	Name    string    `json:"name"`
	IsConst bool      `json:"isConst"`
	Const   *Constant `json:"const"`
	IsVar   bool      `json:"IsVar"`
	Var     *Variable `json:"var"`
	IsType  bool      `json:"isType"`
	Type    *Type     `json:"type"`
	IsFunc  bool      `json:"IsFunc"`
	Func    *Function `json:"func"`
}

func (d *GoDecl) Imports() map[string]string {
	panic("not implemented")
}
