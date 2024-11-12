package generate

import (
	"github.com/mikerybka/pkg/util"
)

type App struct {
	Name        util.Name   `json:"name"`
	Description string      `json:"descsription"`
	LogoSVG     string      `json:"logoSVG"`
	Types       []util.Type `json:"types"`
}
