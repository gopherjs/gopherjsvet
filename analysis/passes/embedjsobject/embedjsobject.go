package embedjsobject

import (
	"go/ast"
	"reflect"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/gopherjs/gopherjsvet/internal"
)

var Analyzer = &analysis.Analyzer{
	Name:     "embedjsobject",
	Doc:      `Whenever js tags are present, a *js.Object field is embedded as the first struct field`,
	URL:      "https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Field)(nil),
	}

	inspector.WithStack(nodeFilter, func(node ast.Node, push bool, stack []ast.Node) bool {
		if !push {
			return true
		}
		field, _ := node.(*ast.Field)
		if field.Tag == nil {
			return true
		}
		tv, _ := strconv.Unquote(field.Tag.Value)
		if _, ok := reflect.StructTag(tv).Lookup("js"); !ok {
			return true
		}
		fieldList := findFieldList(stack)
		for i, field := range fieldList.List {
			if len(field.Names) > 0 {
				// Not embedded
				continue
			}
			star, ok := field.Type.(*ast.StarExpr)
			if !ok {
				// The jsobjectptr analyzer will complain if this is a
				// non-pointer to js.Object, so we can continue here.
				continue
			}
			if internal.Is_jsObject(pass, star.X) {
				if i != 0 {
					pass.Reportf(star.Pos(), "js.Object must be first struct field")
				}
			}
		}
		return true
	})
	return nil, nil
}

func findFieldList(stack []ast.Node) *ast.FieldList {
	for i := len(stack) - 2; i > 0; i-- {
		if fl, ok := stack[i].(*ast.FieldList); ok {
			return fl
		}
	}
	panic("no struct declaration found")
}
