package golang

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"unicode"

	"github.com/mikerybka/pkg/util"
)

type File struct {
	Imports util.ImportMap
	Decls   []Decl
}

func (f *File) AddType(t *util.Type) {
	// Update imports
	newImports := t.Imports()
	for k, v := range newImports {
		if f.Imports == nil {
			f.Imports = util.ImportMap{}
		}
		f.Imports[k] = v
	}

	// Add type declaration
	f.Decls = append(f.Decls, Decl{
		IsType: true,
		Type:   t,
	})

	// Add method declarations
	for _, m := range t.Methods {
		typ := t.Name.GoExported()
		if t.IsStruct {
			typ = "*" + typ
		}
		f.Decls = append(f.Decls, Decl{
			IsMethod: true,
			Method: &util.Method{
				Recv: &util.Field{
					Name: util.NewName(t.Name.LocalVarName()),
					Type: typ,
				},
				Fn: &m,
			},
		})
	}
}

func (f *File) Write(path string) error {
	// Create the file
	osfile, err := os.Create(path)
	if err != nil {
		return err
	}

	// Write the package line
	pkgname := filepath.Base(filepath.Dir(path))
	_, err = fmt.Fprintf(osfile, "package %s\n\n", pkgname)
	if err != nil {
		return err
	}

	// Write imports
	if len(f.Imports) > 0 {
		_, err = fmt.Fprintf(osfile, "import (\n")
		if err != nil {
			return err
		}

		for name, from := range f.Imports {
			_, err = fmt.Fprintf(osfile, "\t")
			if err != nil {
				return err
			}

			if name != filepath.Base(from) {
				_, err = fmt.Fprintf(osfile, "%s ", name)
				if err != nil {
					return err
				}
			}

			_, err = fmt.Fprintf(osfile, "\"%s\"\n", from)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Fprintf(osfile, ")\n\n")
		if err != nil {
			return err
		}
	}

	// Write decls
	for _, decl := range f.Decls {
		if decl.IsType {
			t := decl.Type

			_, err = fmt.Fprintf(osfile, "type %s", t.Name.GoExported())
			if err != nil {
				return err
			}

			if t.IsScalar {
				_, err = fmt.Fprintf(osfile, " %s\n", t.ElemType)
				if err != nil {
					return err
				}
			} else if t.IsArray {
				_, err = fmt.Fprintf(osfile, " []%s\n", t.ElemType)
				if err != nil {
					return err
				}
			} else if t.IsMap {
				_, err = fmt.Fprintf(osfile, " map[string]%s\n", t.ElemType)
				if err != nil {
					return err
				}
			} else if t.IsStruct {
				_, err = fmt.Fprintf(osfile, " struct {\n")
				if err != nil {
					return err
				}

				for _, f := range t.Fields {
					name := f.Name.GoExported()
					_, err = fmt.Fprintf(osfile, "\t%s %s `json:\"%s\"`\n", name, f.Type, lowerFirstLetter(name))
					if err != nil {
						return err
					}
				}

				_, err = fmt.Fprintf(osfile, "}\n")
				if err != nil {
					return err
				}
			}
		}

		if decl.IsMethod {
			m := decl.Method

			_, err = fmt.Fprintf(osfile, "func (%s %s) %s(",
				m.Recv.Name.LocalVarName(),
				m.Recv.Type,
				m.Fn.Name.GoExported(),
			)
			if err != nil {
				return err
			}

			for i, f := range m.Fn.Inputs {
				if i > 0 {
					_, err = fmt.Fprintf(osfile, ", ")
					if err != nil {
						return err
					}
				}

				_, err = fmt.Fprintf(osfile, "%s %s", f.Name.GoUnexported(), f.Type)
				if err != nil {
					return err
				}
			}

			_, err = fmt.Fprintf(osfile, ")")
			if err != nil {
				return err
			}

			if len(m.Fn.Outputs) > 1 {
				_, err = fmt.Fprintf(osfile, " (")
				if err != nil {
					return err
				}
			}

			for i, f := range m.Fn.Outputs {
				if i > 0 {
					_, err = fmt.Fprintf(osfile, ", ")
					if err != nil {
						return err
					}
				}

				_, err = fmt.Fprintf(osfile, f.Type)
				if err != nil {
					return err
				}
			}

			if len(m.Fn.Outputs) > 1 {
				_, err = fmt.Fprintf(osfile, ")")
				if err != nil {
					return err
				}
			}

			_, err = fmt.Fprintf(osfile, "{\n")
			if err != nil {
				return err
			}

			for _, st := range m.Fn.Body {
				if st.IsReturn {
					_, err = fmt.Fprintf(osfile, "\treturn %s\n", st.Return.GoString())
					if err != nil {
						return err
					}
				}
				if st.IsAssign {
					_, err = fmt.Fprintf(osfile, "\t%s := %s\n", st.Name, st.Value.GoString())
					if err != nil {
						return err
					}
				}
				if st.IsIf {
					_, err = fmt.Fprintf(osfile, "\t%s := %s\n", st.Name, st.Value.GoString())
					if err != nil {
						return err
					}
				}
			}

			_, err = fmt.Fprintf(osfile, "}\n")
			if err != nil {
				return err
			}
		}
	}

	// go fmt
	cmd := exec.Command("go", "fmt", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		panic(err)
	}

	return nil
}

func lowerFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
