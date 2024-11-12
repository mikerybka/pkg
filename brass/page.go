package brass

type Page struct {
	Title    string   `json:"title"`
	Actions  []Action `json:"actions"`
	DataType *Type    `json:"dataType"`
	Data     any      `json:"data"`
}
