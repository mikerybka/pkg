package brass

import "github.com/mikerybka/pkg/english"

type Method struct {
	Name    english.Name `json:"name"`
	Inputs  []Field      `json:"inputs"`
	Outputs []Field      `json:"outputs"`
	Body    []Statement  `json:"body"`
}
