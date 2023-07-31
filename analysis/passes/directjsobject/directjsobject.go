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
		detectRawJSObject(pass, node)
		return true
	}
	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}

func detectRawJSObject(pass *analysis.Pass, node ast.Node) {
	if node == nil {
		return
	}
	switch t := node.(type) {
	case *ast.ValueSpec:
		if val, ok := derefPointer(t.Type); ok {
			objMustBeEmbedded(pass, node, val)
		}
	case *ast.ArrayType:
		if arrTypeExpr, ok := derefPointer(t.Elt); ok {
			objMustBeEmbedded(pass, node, arrTypeExpr)
		}
	case *ast.Field:
		switch ft := t.Type.(type) {
		case *ast.SelectorExpr:
			objMustBeEmbedded(pass, node, ft)
		}
	case *ast.MapType:
		if valExpr, ok := derefPointer(t.Value); ok {
			objMustBeEmbedded(pass, node, valExpr)
		}
		if keyExpr, ok := derefPointer(t.Key); ok {
			objMustBeEmbedded(pass, node, keyExpr)
		}
		return
	case *ast.CompositeLit:
		var expr *ast.SelectorExpr
		switch et := t.Type.(type) {
		case *ast.SelectorExpr:
			expr = et
		default:
			return
		}
		objMustBeEmbedded(pass, node, expr)
		return
	default:
		// buf := &bytes.Buffer{}
		// printer.Fprint(buf, pass.Fset, node)
		// fmt.Printf("%T: %s\n", node, buf.String())
	}
}

func derefPointer(node ast.Node) (*ast.SelectorExpr, bool) {
	switch x := node.(type) {
	case *ast.SelectorExpr:
		return x, true
	case *ast.StarExpr:
		child, ok := x.X.(*ast.SelectorExpr)
		if ok {
			return child, true
		}
	}
	return nil, false
}

func objMustBeEmbedded(pass *analysis.Pass, node ast.Node, expr *ast.SelectorExpr) {
	if expr.Sel.Name != "Object" {
		return
	}
	obj := pass.TypesInfo.ObjectOf(expr.Sel)
	if obj == nil {
		return
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return
	}
	if pkg.Path() != "github.com/gopherjs/gopherjs/js" {
		return
	}

	pass.Reportf(node.Pos(), "js.Object must be embedded in a struct")
}
