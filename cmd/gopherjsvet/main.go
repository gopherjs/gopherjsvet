package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/gopherjs/gopherjsvet/analysis/passes/directjsobject"
	"github.com/gopherjs/gopherjsvet/analysis/passes/jsobjectptr"
)

func main() {
	multichecker.Main(
		jsobjectptr.Analyzer,
		directjsobject.Analyzer,
	)
}
