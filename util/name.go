package util

import (
	"strings"
	"unicode"
)

func NewName(s string) Name {
	words := strings.Split(s, " ")
	name := Name{}
	for _, w := range words {
		if w != "" {
			name = append(name, Word(w))
		}
	}
	return name
}

type Name []Word

func (n Name) LocalVarName() string {
	return string(n[0][0])
}

func (n Name) GoUnexported() string {
	s := n.GoExported()
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func (n Name) String() string {
	s := ""
	for i, w := range n {
		if i > 0 {
			s += " "
		}
		s += w.String()
	}
	return s
}

// SnakeCase returns the pascal case representation of the string.
// Ex: "Green Button" => "green_button"
func (n Name) SnakeCase() string {
	s := ""
	for i, w := range n {
		if i > 0 {
			s += "_"
		}
		s += w.StripNonAlphaNumeric().Lower().String()
	}
	return s
}

// ID returns the id friendly string.
// Ex: "Green Button" => "green_button"
func (n Name) ID() string {
	return n.SnakeCase()
}

// GoExported returns an exported Go name.
// Ex: "Green Button" => "GreenButton"
func (n Name) GoExported() string {
	s := ""
	for _, w := range n {
		s += w.StripNonAlphaNumeric().Title().String()
	}
	return s
}
