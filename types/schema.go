package types

type Schema struct {
	ID     string  `json:"id"`
	Desc   string  `json:"desc"`
	Fields []Field `json:"fields"`
}
