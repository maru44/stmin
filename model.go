package stool

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"
)

type (
	ExprType string

	TypePrefix struct {
		IsSlice bool
		IsArray bool
		IsPtr   bool

		Len int
	}

	TypeInfo struct {
		Expr ast.Expr

		ExprType ExprType
		Prefixes []TypePrefix
	}
)

const (
	ExprTypeIdent         = ExprType("ident")
	ExprTypeStructType    = ExprType("struct_type")
	ExprTypeFuncType      = ExprType("func_type")
	ExprTypeInterfaceType = ExprType("interface_type")
	ExprTypeMapType       = ExprType("map_type")
	ExprTypeSelectorExpr  = ExprType("selector_expr")
	ExprTypeStarExpr      = ExprType("star_expr")
	ExprTypeBinaryExpr    = ExprType("binary_expr")
)

func (t *TypePrefix) String() string {
	switch {
	case t.IsArray:
		return fmt.Sprintf("[%d]", t.Len)
	case t.IsSlice:
		return "[]"
	case t.IsPtr:
		return "*"
	}
	panic("must not reach here ")
}

func (t *TypePrefix) IsValid() bool {
	return !t.IsArray && !t.IsPtr && !t.IsSlice
}

func TypePrefixFromString(s string) (out TypePrefix) {
	switch s {
	case "*":
		out.IsPtr = true
	case "[]":
		out.IsSlice = true
	default:
		if !strings.HasPrefix(s, "[") || !strings.HasSuffix(s, "]") {
			return
		}
		l, err := strconv.Atoi(strings.TrimLeft(strings.TrimRight(s, "]"), "["))
		if err != nil {
			return
		}
		out.IsArray = true
		out.Len = l
	}
	return
}
