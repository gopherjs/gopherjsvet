package jsobjectptr

import (
	"fmt"
	"go/ast"

	"github.com/gopherjs/gopherjsvet/internal"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "jsobjectptr",
	Doc:      `js.Object must always be a pointer`,
	URL:      "https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.SelectorExpr)(nil),
	}

	inspector.WithStack(nodeFilter, func(node ast.Node, push bool, stack []ast.Node) bool {
		internal.Dump(pass, node)
		parent := stack[len(stack)-2]
		internal.Dump(pass, parent)
		if !push {
			return true
		}
		if !internal.Is_jsObject(pass, node) {
			fmt.Printf("not js.Object [%T]\n", node)
			return true
		}
		if _, ok := parent.(*ast.StarExpr); !ok {
			pass.Reportf(parent.Pos(), "js.Object must always be a pointer")
		}
		return true
	})
	return nil, nil
}
