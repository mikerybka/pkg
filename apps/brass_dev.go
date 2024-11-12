package apps

import "github.com/mikerybka/pkg/util"

var BrassDev = util.App{
	Name: util.NewName("Brass Dev"),
	Types: []util.Type{
		{
			Name:        util.NewName("Type"),
			Description: "",
			IsStruct:    true,
			Fields: []util.Field{
				{
					Name:        util.NewName("Name"),
					Description: "",
					Type:        "util.Name",
				},
				{
					Name:        util.NewName("Description"),
					Description: "",
					Type:        "string",
				},
				{
					Name:        util.NewName("Is Scalar"),
					Description: "",
					Type:        "bool",
				},
				{
					Name:        util.NewName("Kind"),
					Description: "",
					Type:        "string",
				},
				{
					Name:        util.NewName("Is Array"),
					Description: "",
					Type:        "bool",
				},
				{
					Name:        util.NewName("Is Map"),
					Description: "",
					Type:        "bool",
				},
				{
					Name:        util.NewName("Elem Type"),
					Description: "",
					Type:        "string",
				},
				{
					Name:        util.NewName("Is Struct"),
					Description: "",
					Type:        "bool",
				},
				{
					Name:        util.NewName("Fields"),
					Description: "",
					Type:        "[]util.Field",
				},
				{
					Name:        util.NewName("Methods"),
					Description: "",
					Type:        "map[string]util.Function",
				},
				{
					Name:        util.NewName("Default JSON"),
					Description: "",
					Type:        "string",
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
					Type:        "util.Name",
				},
				{
					Name:        util.NewName("Description"),
					Description: "",
					Type:        "string",
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
