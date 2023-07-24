package directjsobject

import (
	"github.com/gopherjs/gopherjs/js"
)

var x = js.Object{} // want "js.Object must be embedded in a struct"

var y = []js.Object{} // want "js.Object must be embedded in a struct"

var _ = [10]js.Object{} // want "js.Object must be embedded in a struct"

var _ map[string]js.Object // want "js.Object must be embedded in a struct"
