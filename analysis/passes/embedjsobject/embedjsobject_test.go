package embedjsobject_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/gopherjs/gopherjsvet/analysis/passes/embedjsobject"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, embedjsobject.Analyzer, "embedjsobject")
}
