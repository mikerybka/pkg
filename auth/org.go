package auth

import "github.com/mikerybka/pkg/util"

type Org struct {
	ID      string
	Members util.Set[string]
}

func (o *Org) CanRead(userID string) bool {
	if userID == "" {
		return false
	}
	return o.Members.Has(userID)
}

func (o *Org) CanWrite(userID string) bool {
	if userID == "" {
		return false
	}
	return o.Members.Has(userID)
}
