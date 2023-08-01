package internal

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Returns true if expr represent a js.Object or *js.Object.
func Is_jsObject(pass *analysis.Pass, node ast.Node) bool {
	expr, ok := node.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	if expr.Sel.Name != "Object" {
		return false
	}
	obj := pass.TypesInfo.ObjectOf(expr.Sel)
	if obj == nil {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return pkg.Path() == "github.com/gopherjs/gopherjs/js"
}
