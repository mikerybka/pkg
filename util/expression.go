package util

import (
	"fmt"
	"strings"
)

type Expression struct {
	IsLiteral bool   `json:"isLiteral"`
	Value     string `json:"value"`

	IsCall bool          `json:"isCall"`
	Fn     string        `json:"fn"`
	Args   []*Expression `json:"args"`

	IsRef bool   `json:"isRef"`
	Ref   string `json:"ref"`
}

func (e *Expression) Imports() ImportMap {
	imports := ImportMap{}

	if e.IsRef {
		from, _ := parseName(e.Ref)
		if from != "" {
			imports[from] = importPath(from)
		}
	}

	if e.IsCall {
		from, _ := parseName(e.Fn)
		if from != "" {
			imports[from] = importPath(from)
		}

		for _, arg := range e.Args {
			for k, v := range arg.Imports() {
				imports[k] = v
			}
		}
	}

	return imports
}

func (e *Expression) GoString() string {
	if e.IsLiteral {
		return e.Value
	}
	if e.IsCall {
		if e.Fn == "!" {
			return fmt.Sprintf("!%s", e.Args[0].GoString())
		}
		if e.Fn == "+" {
			return fmt.Sprintf("%s + %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "-" {
			return fmt.Sprintf("%s - %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "*" {
			return fmt.Sprintf("%s * %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "/" {
			return fmt.Sprintf("%s / %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "%" {
			return fmt.Sprintf("%s % %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "==" {
			return fmt.Sprintf("%s == %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "!=" {
			return fmt.Sprintf("%s != %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "<" {
			return fmt.Sprintf("%s < %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "<=" {
			return fmt.Sprintf("%s <= %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == ">" {
			return fmt.Sprintf("%s > %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == ">=" {
			return fmt.Sprintf("%s >= %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "&&" {
			return fmt.Sprintf("%s && %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "||" {
			return fmt.Sprintf("%s || %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "&" {
			return fmt.Sprintf("%s & %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "|" {
			return fmt.Sprintf("%s | %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "&^" {
			return fmt.Sprintf("%s &^ %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == "<<" {
			return fmt.Sprintf("%s << %s", e.Args[0].GoString(), e.Args[1].GoString())
		}
		if e.Fn == ">>" {
			return fmt.Sprintf("%s >> %s", e.Args[0].GoString(), e.Args[1].GoString())
		}

		args := []string{}
		for _, arg := range e.Args {
			args = append(args, arg.GoString())
		}
		return fmt.Sprintf("%s(%s)", e.Fn, strings.Join(args, ","))
	}
	if e.IsRef {
		return e.Ref
	}
	panic("unreachable")
}
