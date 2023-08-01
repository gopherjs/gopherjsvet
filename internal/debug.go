package internal

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"os"

	"golang.org/x/tools/go/analysis"
)

func Dump(pass *analysis.Pass, node ast.Node) {
	if os.Getenv("DEBUG") == "" {
		return
	}
	buf := &bytes.Buffer{}
	printer.Fprint(buf, pass.Fset, node)
	fmt.Printf("%T: %s\n", node, buf.String())
}
