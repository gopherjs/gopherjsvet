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
	switch t := node.(type) {
	case *ast.StructType:
		for _, f := range t.Fields.List {
			switch ft := f.Type.(type) {
			case *ast.SelectorExpr:
				objMustBeEmbedded(pass, node, ft)
			}
		}
	case *ast.MapType:
		if valExpr, ok := t.Value.(*ast.SelectorExpr); ok {
			objMustBeEmbedded(pass, node, valExpr)
		}
		keyExpr, ok := t.Key.(*ast.SelectorExpr)
		if !ok {
			return
		}
		objMustBeEmbedded(pass, node, keyExpr)
		return
	case *ast.CompositeLit:
		var expr *ast.SelectorExpr
		switch et := t.Type.(type) {
		case *ast.SelectorExpr:
			expr = et
		case *ast.ArrayType:
		outer:
			for {
				switch at := et.Elt.(type) {
				case *ast.SelectorExpr:
					expr = at
					break outer
				case *ast.ArrayType:
					et = at
				default:
					return
				}
			}
		default:
			return
		}
		objMustBeEmbedded(pass, node, expr)
		return
	}
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
