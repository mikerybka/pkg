package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikerybka/pkg/util"
)

type Data struct {
	Dir string
}

func (d *Data) User(id string) *User {
	path := filepath.Join(d.Dir, "users", id)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	user := &User{}
	json.NewDecoder(f).Decode(user)
	return user
}

func (d *Data) Org(id string) *Org {
	path := filepath.Join(d.Dir, "orgs", id)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	org := &Org{}
	json.NewDecoder(f).Decode(org)
	return org
}

func (d *Data) UserByPhone(phone util.PhoneNumber) (*User, error) {
	path := filepath.Join(d.Dir, "users/phone_index")
	phoneIndex := util.ReadJSONFile[map[util.PhoneNumber]string](path)
	id, ok := phoneIndex[phone]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	path = filepath.Join(d.Dir, "users", id)
	u := util.ReadJSONFile[User](path)
	return &u, nil
}

func (d *Data) SaveUser(u *User) error {
	// Check for conflicting phone numbers
	path := filepath.Join(d.Dir, "users/phone_index")
	phoneIndex := util.ReadJSONFile[map[string]string](path)
	id, ok := phoneIndex[u.Phone]
	if ok && u.ID != "" && u.ID != id {
		return fmt.Errorf("phone registered with another account")
	}

	// Update the index
	if phoneIndex == nil {
		phoneIndex = make(map[string]string)
	}
	phoneIndex[u.Phone] = u.ID
	err := util.WriteJSONFile(path, phoneIndex)
	if err != nil {
		panic(err)
	}

	// Write the user file
	path = filepath.Join(d.Dir, "users", u.ID)
	err = util.WriteJSONFile(path, u)
	if err != nil {
		panic(err)
	}

	return nil
}
