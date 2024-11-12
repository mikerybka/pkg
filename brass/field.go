package brass

import "github.com/mikerybka/pkg/english"

type Field struct {
	Name english.Name `json:"name"`
	Type *Type        `json:"type"`
}

func (f *Field) ID() string {
	return f.Name.ID()
}
