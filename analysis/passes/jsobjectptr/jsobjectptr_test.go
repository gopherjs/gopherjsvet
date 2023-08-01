package jsobjectptr_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/gopherjs/gopherjsvet/analysis/passes/jsobjectptr"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, jsobjectptr.Analyzer, "jsobjectptr", "renamedimport", "otherjsimport")
}
