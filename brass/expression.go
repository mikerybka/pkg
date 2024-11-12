package brass

type Expression struct {
	IsLiteral bool   `json:"isLiteral"`
	Value     string `json:"value"`

	IsCall bool          `json:"isCall"`
	Fn     *Ref          `json:"fn"`
	Args   []*Expression `json:"args"`

	IsRef bool `json:"isRef"`
	Ref   *Ref `json:"ref"`
}

func (e *Expression) Eval(env map[string]string) {
	panic("not implemented")
}
