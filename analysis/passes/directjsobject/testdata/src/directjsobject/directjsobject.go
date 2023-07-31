package directjsobject

import (
	"github.com/gopherjs/gopherjs/js"
)

var x = js.Object{} // want "js.Object must be embedded in a struct"

var _ = &js.Object{} // want "js.Object must be embedded in a struct"

var y = []js.Object{} // want "js.Object must be embedded in a struct"

var _ = []*js.Object{} // want "js.Object must be embedded in a struct"

var _ = [10]js.Object{} // want "js.Object must be embedded in a struct"

var _ = [10]*js.Object{} // want "js.Object must be embedded in a struct"

var _ = map[string]js.Object{} // want "js.Object must be embedded in a struct"

var _ = map[string]*js.Object{} // want "js.Object must be embedded in a struct"

var _ = map[js.Object]string{} // want "js.Object must be embedded in a struct"

var _ = map[*js.Object]string{} // want "js.Object must be embedded in a struct"

var _ = [][]js.Object{} // want "js.Object must be embedded in a struct"

var _ = [][][]js.Object{} // want "js.Object must be embedded in a struct"

var _ = struct{ obj js.Object }{} // want "js.Object must be embedded in a struct"

var _ = struct{ js.Object }{}

var _ = struct {
	sl  []js.Object       // want "js.Object must be embedded in a struct"
	sl2 [][][]js.Object   // want "js.Object must be embedded in a struct"
	m   map[int]js.Object // want "js.Object must be embedded in a struct"
	st  struct {
		j  js.Object    // want "js.Object must be embedded in a struct"
		a5 [5]js.Object // want "js.Object must be embedded in a struct"
	}
}{}
