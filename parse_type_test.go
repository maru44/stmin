package stool

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_purgePointerOrSlice(t *testing.T) {
	tests := []struct {
		name  string
		input ast.Expr

		wantExpr   ast.Expr
		wantPrefix TypePrefix
		wantFinish bool
	}{
		{
			name: "ok: ptr",
			input: &ast.StarExpr{
				X: &ast.Ident{
					Name: "TestIdent",
				},
			},
			wantExpr: &ast.Ident{
				Name: "TestIdent",
			},
			wantPrefix: TypePrefix{IsPtr: true},
		},
		{
			name: "ok: slice",
			input: &ast.ArrayType{
				Elt: &ast.StarExpr{
					X: &ast.Ident{
						Name: "TestIdent2",
					},
				},
			},
			wantExpr: &ast.StarExpr{
				X: &ast.Ident{
					Name: "TestIdent2",
				},
			},
			wantPrefix: TypePrefix{
				IsSlice: true,
			},
		},
		{
			name: "ok: array",
			input: &ast.ArrayType{
				Elt: &ast.StarExpr{
					X: &ast.Ident{
						Name: "TestIdent3",
					},
				},
				Len: &ast.BasicLit{
					Kind: token.Token(222),
				},
			},
			wantExpr: &ast.StarExpr{
				X: &ast.Ident{
					Name: "TestIdent3",
				},
			},
			wantPrefix: TypePrefix{
				IsArray: true,
				Len:     222,
			},
		},
		{
			name:       "not ok",
			input:      &ast.FuncType{},
			wantExpr:   &ast.FuncType{},
			wantPrefix: TypePrefix{},
			wantFinish: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gotExpr, gotPrefix, gotFinish := purgePointerOrSlice(tt.input)
			assert.Equal(t, tt.wantExpr, gotExpr)
			assert.Equal(t, tt.wantPrefix, gotPrefix)
			assert.Equal(t, tt.wantFinish, gotFinish)
		})
	}
}
