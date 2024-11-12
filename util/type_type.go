package util

var TypeType = &Type{
	IsStruct: true,
	Fields: []Field{
		{
			Name: NewName("Is scalar"),
			Desc: "",
			Type: "bool",
		},
		{
			Name: NewName("Kind"),
			Desc: "",
			Type: "string",
		},
		{
			Name: NewName("Is pointer"),
			Desc: "",
			Type: "bool",
		},
		{
			Name: NewName("Is array"),
			Desc: "",
			Type: "bool",
		},
		{
			Name: NewName("Is map"),
			Desc: "",
			Type: "bool",
		},
		{
			Name: NewName("Elem type"),
			Desc: "",
			Type: "string",
		},
		{
			Name: NewName("Is struct"),
			Desc: "",
			Type: "bool",
		},
		{
			Name: NewName("Fields"),
			Desc: "",
			Type: "[]util.Field",
		},
	},
}
