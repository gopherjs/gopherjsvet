package directjsobject_test

import (
	"testing"

	_ "github.com/gopherjs/gopherjs/js"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/assign"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	tests := []string{"directjsobject"}
	analysistest.Run(t, testdata, assign.Analyzer, tests...)
}
