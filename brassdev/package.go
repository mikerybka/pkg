package brassdev

import "github.com/mikerybka/pkg/util"

type Package struct {
	ID        string                   `json:"id"`
	Desc      string                   `json:"desc"`
	Comments  []util.Comment           `json:"comments"`
	Types     map[string]util.Type     `json:"types"`
	Functions map[string]util.Function `json:"functions"`
	Variables map[string]util.Value    `json:"variables"`
	Constants map[string]util.Value    `json:"constants"`
}
