package brass

import "fmt"

func NewGoFile(pkg string) *GoFile {
	return &GoFile{
		Pkg: pkg,
	}
}

type GoFile struct {
	Pkg     string `json:"pkg"`
	Imports map[string]string
	Decls   []GoDecl `json:"decls"`
}

func (f *GoFile) hasName(n string) bool {
	for _, d := range f.Decls {
		if d.Name == n {
			return true
		}
	}
	return false
}

func (f *GoFile) AddDecl(d GoDecl) error {
	if f.hasName(d.Name) {
		return fmt.Errorf("name %s already defined", d.Name)
	}
	f.Decls = append(f.Decls, d)
	return nil
}
