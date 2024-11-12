package generate

import (
	"path/filepath"

	"github.com/mikerybka/pkg/golang"
	"github.com/mikerybka/pkg/util"
)

func GoTypes(types []util.Type, dst string) error {
	for _, t := range types {
		filename := t.Name.ID() + ".go"
		path := filepath.Join(dst, filename)
		f := &golang.File{}
		f.AddType(&t)
		err := f.Write(path)
		if err != nil {
			return err
		}
	}
	return nil
}
