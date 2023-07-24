package directjsobject

import (
	othername "github.com/gopherjs/gopherjs/js"
)

var x = othername.Object{} // want "js.Object must be embedded in a struct"
