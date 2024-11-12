package brass

import (
	"fmt"
	"strings"
)

type Action struct {
	Path   string
	Method string
	Title  string
	Desc   string
	Fields []Field
}

func (a *Action) Validate() error {
	// Check that the path starts with a slash.
	if !strings.HasPrefix(a.Path, "/") {
		return fmt.Errorf("path must start with /")
	}

	// Check existence of method.
	if a.Method == "" {
		return fmt.Errorf("no method")
	}

	// Check existence of title.
	if a.Title == "" {
		return fmt.Errorf("no title")
	}

	// Check for multiple occurneces of the same ID.
	for i, f := range a.Fields {
		for j := i + 1; j < len(a.Fields); j++ {
			if f.ID() == a.Fields[j].ID() {
				return fmt.Errorf("fields %d and %d have the same ID", i, j)
			}
		}
	}

	return nil
}
