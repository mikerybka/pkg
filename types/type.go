package types

type Type struct {
	ID   string `json:"id"`
	Desc string `json:"desc"`

	Kind   string  `json:"kind"`
	Fields []Field `json:"fields"`
}
