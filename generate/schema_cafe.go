package generate

import (
	"github.com/mikerybka/pkg/util"
)

var SchemaCafe = &App{
	Name:        util.NewName("Schema.cafe"),
	Description: "",
	LogoSVG:     "",
	Types: []util.Type{
		{
			Name:        util.NewName("Schema"),
			Description: "",
			IsStruct:    true,
			Fields: []util.Field{
				{
					Name:        util.NewName("Name"),
					Description: "",
					Type:        "name",
				},
				{
					Name:        util.NewName("Description"),
					Description: "",
					Type:        "text",
				},
				{
					Name:        util.NewName("Fields"),
					Description: "",
					Type:        "[]Field",
				},
			},
		},
		{
			Name:        util.NewName("Field"),
			Description: "",
			IsStruct:    true,
			Fields: []util.Field{
				{
					Name:        util.NewName("Name"),
					Description: "",
					Type:        "name",
				},
				{
					Name:        util.NewName("Description"),
					Description: "",
					Type:        "text",
				},
				{
					Name:        util.NewName("Type"),
					Description: "",
					Type:        "string",
				},
			},
		},
	},
}
