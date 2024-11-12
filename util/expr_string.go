package util

import (
	"fmt"
	"go/ast"
)

// ExprString converts an ast.Expr to its string representation.
func ExprString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		// Basic literals: strings, numbers, etc.
		return e.Value
	case *ast.Ident:
		// Identifiers (variable names, types, etc.)
		return e.Name
	case *ast.BinaryExpr:
		// Binary expressions like `a + b`
		return fmt.Sprintf("%s %s %s", ExprString(e.X), e.Op, ExprString(e.Y))
	case *ast.CallExpr:
		// Function or method calls like `foo(a, b)`
		args := ""
		for i, arg := range e.Args {
			if i > 0 {
				args += ", "
			}
			args += ExprString(arg)
		}
		return fmt.Sprintf("%s(%s)", ExprString(e.Fun), args)
	case *ast.ParenExpr:
		// Parenthesized expressions like `(a + b)`
		return fmt.Sprintf("(%s)", ExprString(e.X))
	case *ast.SelectorExpr:
		// Selector expressions like `pkg.Func` or `obj.Field`
		return fmt.Sprintf("%s.%s", ExprString(e.X), e.Sel.Name)
	case *ast.IndexExpr:
		// Index expressions like `arr[i]`
		return fmt.Sprintf("%s[%s]", ExprString(e.X), ExprString(e.Index))
	case *ast.SliceExpr:
		// Slice expressions like `arr[start:end]`
		low := ""
		high := ""
		if e.Low != nil {
			low = ExprString(e.Low)
		}
		if e.High != nil {
			high = ExprString(e.High)
		}
		return fmt.Sprintf("%s[%s:%s]", ExprString(e.X), low, high)
	case *ast.StarExpr:
		// Star expressions like `*T`
		return fmt.Sprintf("*%s", ExprString(e.X))
	case *ast.UnaryExpr:
		// Unary expressions like `-a`, `*p`, `&x`
		return fmt.Sprintf("%s%s", e.Op, ExprString(e.X))
	case *ast.ArrayType:
		// Array or slice types like `[]T` or `[n]T`
		length := ""
		if e.Len != nil {
			length = ExprString(e.Len)
		}
		return fmt.Sprintf("[%s]%s", length, ExprString(e.Elt))
	case *ast.MapType:
		// Map types like `map[K]V`
		return fmt.Sprintf("map[%s]%s", ExprString(e.Key), ExprString(e.Value))
	case *ast.FuncType:
		// Function types like `func(int, int) string`
		params := ""
		for i, param := range e.Params.List {
			for j, name := range param.Names {
				if i > 0 || j > 0 {
					params += ", "
				}
				params += ExprString(name)
			}
			params += " " + ExprString(param.Type)
		}
		results := ""
		if e.Results != nil {
			for i, result := range e.Results.List {
				if i > 0 {
					results += ", "
				}
				results += ExprString(result.Type)
			}
		}
		if results == "" {
			return fmt.Sprintf("func(%s)", params)
		}
		return fmt.Sprintf("func(%s) %s", params, results)
	case *ast.InterfaceType:
		// Interface types like `interface { ... }`
		return "interface { ... }"
	case *ast.StructType:
		// Struct types like `struct { ... }`
		return "struct { ... }"
	case *ast.Ellipsis:
		// Variadic parameter types like `...T`
		return fmt.Sprintf("...%s", ExprString(e.Elt))
	case *ast.TypeAssertExpr:
		// Type assertions like `x.(T)`
		return fmt.Sprintf("%s.(%s)", ExprString(e.X), ExprString(e.Type))
	case *ast.ChanType:
		// Channel types like `chan T`, `<-chan T`, `chan<- T`
		dir := ""
		if e.Dir == ast.RECV {
			dir = "<-"
		} else if e.Dir == ast.SEND {
			dir = "chan<-"
		}
		return fmt.Sprintf("%schan %s", dir, ExprString(e.Value))
	default:
		// For anything else, we return a placeholder
		return "<unknown>"
	}
}
