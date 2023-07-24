package directjsobject

import (
	"fmt"
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
			sel, ok := t.Type.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if sel.Sel.Name != "Object" {
				return true
			}
			x, ok := sel.X.(*ast.Ident)
			if !ok {
				return true
			}
			if x.Name != "js" {
				return true
			}
			pass.Reportf(node.Pos(), "js.Object must be embedded in a struct")
		case *ast.SelectorExpr:
			fmt.Printf("XXX: %v\n", t.X)
		default:
			fmt.Printf("%v %T\n", node, node)

		}
		return true
	}
	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
