//go:build exclude
// +build exclude

package main

import (
	"go/ast"

	"github.com/k0kubun/pp"
	"github.com/maru44/stool"
	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}
	ps, err := packages.Load(cfg, "github.com/maru44/stool/_example/data")
	if err != nil {
		panic(err)
	}

	for _, p := range ps {
		for _, f := range p.Syntax {
			for _, decl := range f.Decls {
				if it, ok := decl.(*ast.GenDecl); ok {
					for _, spec := range it.Specs {
						switch ts := spec.(type) {
						case *ast.TypeSpec:
							pp.Println(stool.PurgePrefixes(ts.Type))
						}
					}
				}
			}
		}
	}
}

/* result like this

stool.TypeInfo{
  Expr: &ast.Ident{
    NamePos: 32,
    Name:    "int",
    Obj:     (*ast.Object)(nil),
  },
  ExprType: "ident",
  Prefixes: []stool.TypePrefix{
    stool.TypePrefix{
      IsSlice: true,
      IsArray: false,
      IsPtr:   false,
      Len:     0,
    },
  },
}

stool.TypeInfo{
  Expr: &ast.Ident{
    NamePos: 45,
    Name:    "string",
    Obj:     (*ast.Object)(nil),
  },
  ExprType: "ident",
  Prefixes: []stool.TypePrefix{
    stool.TypePrefix{
      IsSlice: false,
      IsArray: false,
      IsPtr:   true,
      Len:     0,
    },
  },
}

stool.TypeInfo{
  Expr: &ast.MapType{
    Map: 71,
    Key: &ast.Ident{
      NamePos: 75,
      Name:    "string",
      Obj:     (*ast.Object)(nil),
    },
    Value: &ast.StarExpr{
      Star: 82,
      X:    &ast.StructType{
        Struct: 83,
        Fields: &ast.FieldList{
          Opening: 89,
          List:    []*ast.Field{},
          Closing: 90,
        },
        Incomplete: false,
      },
    },
  },
  ExprType: "map_type",
  Prefixes: []stool.TypePrefix{
    stool.TypePrefix{
      IsSlice: true,
      IsArray: false,
      IsPtr:   false,
      Len:     0,
    },
    stool.TypePrefix{
      IsSlice: false,
      IsArray: false,
      IsPtr:   true,
      Len:     0,
    },
    stool.TypePrefix{
      IsSlice: true,
      IsArray: false,
      IsPtr:   false,
      Len:     0,
    },
    stool.TypePrefix{
      IsSlice: false,
      IsArray: true,
      IsPtr:   false,
      Len:     5,
    },
  },
}

*/
