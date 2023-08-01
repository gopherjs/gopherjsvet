package jsobjectptr_test

import (
	"testing"

	"github.com/gopherjs/gopherjsvet/analysis/passes/jsobjectptr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, jsobjectptr.Analyzer, "jsobjectptr")
}
