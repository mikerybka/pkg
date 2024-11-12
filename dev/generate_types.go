package dev

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mikerybka/pkg/golang"
	"github.com/mikerybka/pkg/util"
)

func GenerateTypes() error {
	typeIDs, err := ListTypeIDs()
	if err != nil {
		return fmt.Errorf("listing types: %s", err)
	}

	for _, typeID := range typeIDs {
		typ, err := GetType(typeID)
		if err != nil {
			return fmt.Errorf("getting type %s: %s", typeID, err)
		}
		snakecase := strings.ReplaceAll(typeID, "-", "_")
		path := filepath.Join(
			util.HomeDir(),
			fmt.Sprintf("src/main/pkg/types/%s.go", snakecase),
		)
		err = golang.WriteTypeFile(path, typ)
		if err != nil {
			return fmt.Errorf("writing file %s: %s", path, err)
		}
	}

	return nil
}
