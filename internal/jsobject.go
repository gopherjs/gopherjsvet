package internal

import (
	"go/ast"
	"go/types"

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
	// For embedded structs, ObjectOf returns an instance of the struct field,
	// not the type itself. So this does the necessary conversion.
	if varObj, ok := obj.(*types.Var); ok {
		switch typ := varObj.Type().(type) {
		case *types.Named:
			obj = typ.Obj()
		case *types.Pointer:
			named, ok := typ.Elem().(*types.Named)
			if ok {
				obj = named.Obj()
			} else {
				return false
			}
		default:
			return false
		}
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}
	return pkg.Path() == "github.com/gopherjs/gopherjs/js"
}
