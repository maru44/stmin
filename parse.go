package stmin

import (
	"go/ast"
)

func GetTypeInfo(input ast.Expr) TypeInfo {
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
	}

	// return ex
	switch ex.(type) {
	case *ast.Ident:
		out.ExprType = ExprTypeIdent
	case *ast.StructType:
		out.ExprType = ExprTypeStructType
	case *ast.FuncType:
		out.ExprType = ExprTypeFuncType
	case *ast.InterfaceType:
		out.ExprType = ExprTypeInterfaceType
	case *ast.MapType:
		out.ExprType = ExprTypeMapType
	case *ast.SelectorExpr:
		out.ExprType = ExprTypeSelectorExpr
	case *ast.StarExpr:
		out.ExprType = ExprTypeStarExpr
	case *ast.BinaryExpr:
		out.ExprType = ExprTypeBinaryExpr
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
