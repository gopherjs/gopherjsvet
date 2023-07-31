package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/gopherjs/gopherjsvet/analysis/passes/directjsobject"
)

func main() {
	multichecker.Main(directjsobject.Analyzer)
}
