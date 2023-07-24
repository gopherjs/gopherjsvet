package directjsobject

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "directjsobect",
	Doc:      `Do not use items or fields of type js.Object directly`,
	URL:      "https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		switch t := node.(type) {
		case *ast.CompositeLit:
			var expr *ast.SelectorExpr
			switch et := t.Type.(type) {
			case *ast.SelectorExpr:
				expr = et
			case *ast.ArrayType:
				var ok bool
				expr, ok = et.Elt.(*ast.SelectorExpr)
				if !ok {
					return true
				}
			}
			if expr.Sel.Name != "Object" {
				return true
			}
			obj := pass.TypesInfo.ObjectOf(expr.Sel)
			if obj == nil {
				return true
			}
			pkg := obj.Pkg()
			if pkg == nil {
				return true
			}
			if pkg.Path() != "github.com/gopherjs/gopherjs/js" {
				return true
			}

			pass.Reportf(node.Pos(), "js.Object must be embedded in a struct")
		}
		return true
	}
	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
