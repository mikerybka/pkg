package golang

import "github.com/mikerybka/pkg/util"

type Decl struct {
	IsType   bool
	Type     *util.Type
	IsMethod bool
	Method   *util.Method
}
