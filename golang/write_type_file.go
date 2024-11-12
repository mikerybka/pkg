package golang

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/mikerybka/pkg/types"
)

func WriteTypeFile(path string, typ *types.Type) error {
	b, err := templates.ReadFile("templates/type.go.tmpl")
	if err != nil {
		panic(err)
	}
	t := template.New("type").Funcs(template.FuncMap{
		"eq":         eq,
		"pascalCase": pascalCase,
		"snakeCase":  snakeCase,
		"golang":     golang,
	})
	t = template.Must(t.Parse(string(b)))
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	err = t.Execute(f, typ)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("go", "fmt", path)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", out)
		return err
	}
	return nil
}

var acronyms = map[string]string{
	"id":  "ID",
	"ids": "IDs",
}

func eq(a, b string) bool {
	return a == b
}
func pascalCase(s string) string {
	words := strings.Split(s, "-")
	for i, word := range words {
		acronym, ok := acronyms[word]
		if ok {
			words[i] = acronym
		} else {
			words[i] = strings.Title(word)
		}
	}
	return strings.Join(words, "")
}

func snakeCase(s string) string {
	return strings.ReplaceAll(s, "-", "_")
}

func golang(t string) string {
	prefix := "[]"
	if strings.HasPrefix(t, prefix) {
		woPrefix := strings.TrimPrefix(t, prefix)
		if woPrefix == "text" {
			woPrefix = "string"
		}
		if isBuiltin(woPrefix) {
			return t
		}
		return prefix + pascalCase(woPrefix)
	}
	prefix = "map[string]"
	if strings.HasPrefix(t, prefix) {
		woPrefix := strings.TrimPrefix(t, prefix)
		if woPrefix == "text" {
			woPrefix = "string"
		}
		if isBuiltin(woPrefix) {
			return t
		}
		return prefix + pascalCase(woPrefix)
	}
	if t == "text" {
		t = "string"
	}
	if isBuiltin(t) {
		return t
	}
	return pascalCase(t)
}

func isBuiltin(t string) bool {
	return t == "string" || t == "bool" || t == "int64" || t == "float64"
}
