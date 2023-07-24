package directjsobject_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/gopherjs/gopherjsvet/analysis/passes/directjsobject"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, directjsobject.Analyzer, "directjsobject", "otherjsimport", "renamedimport")
}
