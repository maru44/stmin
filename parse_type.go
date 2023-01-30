package stool

import (
	"go/ast"
)

func PurgePrefixes(input ast.Expr) TypeInfo {
	var fin bool
	var prefixes []TypePrefix
	ex := input
	for !fin {
		var pref TypePrefix
		ex, pref, fin = purgePointerOrSlice(ex)
		if fin {
			break
		}
		prefixes = append(prefixes, pref)
	}

	out := TypeInfo{
		Expr:     ex,
		Prefixes: prefixes,
		ExprType: exprType(ex),
	}
	return out
}

func purgePointerOrSlice(ex ast.Expr) (ast.Expr, TypePrefix, bool) {
	switch typ := ex.(type) {
	case *ast.StarExpr:
		return typ.X, TypePrefixFromString("*"), false
	case *ast.ArrayType:
		if typ.Len != nil {
			if b, ok := typ.Len.(*ast.BasicLit); ok {
				return typ.Elt, TypePrefix{
					IsArray: true,
					Len:     int(b.Kind),
				}, false
			}
		}
		return typ.Elt, TypePrefixFromString("[]"), false
	}
	return ex, TypePrefix{}, true
}

func exprType(ex ast.Expr) ExprType {
	switch ex.(type) {
	case *ast.Ident:
		return ExprTypeIdent
	case *ast.StructType:
		return ExprTypeStructType
	case *ast.FuncType:
		return ExprTypeFuncType
	case *ast.InterfaceType:
		return ExprTypeInterfaceType
	case *ast.MapType:
		return ExprTypeMapType
	case *ast.SelectorExpr:
		return ExprTypeSelectorExpr
	case *ast.StarExpr:
		return ExprTypeStarExpr
	case *ast.BinaryExpr:
		return ExprTypeBinaryExpr
	}
	return ExprType("")
}
