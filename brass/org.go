package brass

import "github.com/mikerybka/pkg/util"

type Org struct {
	ID      string
	Members util.Set[string]
}

func (o *Org) CanRead(userID, path string) bool {
	return true
}

func (o *Org) CanWrite(userID, path string) bool {
	if userID == "" {
		return false
	}
	return o.Members.Has(userID)
}
