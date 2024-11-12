package types

import "github.com/mikerybka/pkg/brass"

var PhoneNumber = &brass.Type{
	IsScalar: true,
	Kind:     "string",
}
