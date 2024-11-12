package brass

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/mikerybka/pkg/util"
)

type Type struct {
	IsScalar bool   `json:"isScalar"`
	Kind     string `json:"kind"`

	IsPointer bool  `json:"isPointer"`
	IsArray   bool  `json:"isArray"`
	IsMap     bool  `json:"isMap"`
	ElemType  *Type `json:"elemType"`

	IsStruct bool    `json:"isStruct"`
	Fields   []Field `json:"fields"`

	DefaultJSON string `json:"defaultJSON"`

	Methods map[string]*Function `json:"methods"`
}

func (t *Type) New() any {
	if t.IsScalar {
		def, ok := builtinDefaults[t.Kind]
		if !ok {
			panic("unkown type " + t.Kind)
		}
		return def
	}
	if t.IsPointer {
		return reflect.ValueOf(t.ElemType.New()).Addr().Interface()
	}
	if t.IsArray {
		typ := reflect.ArrayOf(0, reflect.TypeOf(t.ElemType.New()))
		return reflect.New(typ).Elem().Interface()
	}
	if t.IsMap {
		keyT := reflect.TypeOf("")
		valT := reflect.TypeOf(t.ElemType.New())
		typ := reflect.MapOf(keyT, valT)
		return reflect.New(typ).Elem().Interface()
	}
	if t.IsStruct {
		fields := []reflect.StructField{}
		for _, f := range t.Fields {
			fields = append(fields, reflect.StructField{
				Name: f.Name.GoExported(),
				Tag:  reflect.StructTag(fmt.Sprintf(`json:"%s"`, f.Name.ID())),
				Type: reflect.TypeOf(f.Type.New()),
			})
		}
		typ := reflect.StructOf(fields)
		v := reflect.New(typ).Elem().Interface()
		json.Unmarshal([]byte(t.DefaultJSON), &v)
		return v
	}
	return nil
}

func (t *Type) GoString(indent int) string {
	if t.IsScalar {
		return t.Kind
	}
	if t.IsPointer {
		return "*" + t.ElemType.GoString(0)
	}
	if t.IsArray {
		return "[]" + t.ElemType.GoString(0)
	}
	if t.IsMap {
		return "map[string]" + t.ElemType.GoString(0)
	}
	if t.IsStruct {
		s := strings.Builder{}
		s.WriteString("struct {\n")
		for _, f := range t.Fields {
			for i := 0; i < indent+1; i++ {
				s.WriteString("\t")
			}
			s.WriteString(f.Name.GoExported())
			s.WriteString(" ")
			s.WriteString(f.Type.GoString(0))
			s.WriteString("\n")
		}
		s.WriteString("}")
		return s.String()
	}
	panic("unknown type")
}

func (t *Type) WriteGoFile(w io.Writer, pkg, name string) (int, error) {
	return fmt.Fprintf(w, "package %s\n\ntype %s %s\n", pkg, name, t.GoString(0))
}

func (t *Type) GoFile(pkg, typ string) string {
	return fmt.Sprintf("package %s\n\ntype %s %s\n", pkg, typ, t.GoString(0))
}

func (t *Type) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		util.WriteJSON(w, Value{
			Type: "Type",
			Data: t,
		})
		return
	}

	if r.Method == "POST" {
		// Parse
		req := util.ReadJSON[util.MethodCall](r.Body)

		// Execute
		res, err := util.CallMethod(t, req.Method, req.Args)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Respond
		util.WriteJSON(w, res)
		return
	}
}
