package directjsobject

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	_ = js.Object{}      // want "js.Object must be embedded in a struct"
	_ = &js.Object{}     // want "js.Object must be embedded in a struct"
	_ = []js.Object{}    // want "js.Object must be embedded in a struct"
	_ = []*js.Object{}   // want "js.Object must be embedded in a struct"
	_ = [10]js.Object{}  // want "js.Object must be embedded in a struct"
	_ = [10]*js.Object{} // want "js.Object must be embedded in a struct"
)

var (
	_ = map[string]js.Object{}  // want "js.Object must be embedded in a struct"
	_ = map[string]*js.Object{} // want "js.Object must be embedded in a struct"
	_ = map[js.Object]string{}  // want "js.Object must be embedded in a struct"
	_ = map[*js.Object]string{} // want "js.Object must be embedded in a struct"
)

var (
	_ = [][]js.Object{}           // want "js.Object must be embedded in a struct"
	_ = [][][]js.Object{}         // want "js.Object must be embedded in a struct"
	_ = struct{ obj js.Object }{} // want "js.Object must be embedded in a struct"
	_ = struct{ js.Object }{}
)

var _ = struct {
	sl  []js.Object       // want "js.Object must be embedded in a struct"
	sl2 [][][]js.Object   // want "js.Object must be embedded in a struct"
	m   map[int]js.Object // want "js.Object must be embedded in a struct"
	st  struct {
		j  js.Object    // want "js.Object must be embedded in a struct"
		a5 [5]js.Object // want "js.Object must be embedded in a struct"
	}
}{}

var (
	_ js.Object             // want "js.Object must be embedded in a struct"
	_ *js.Object            // want "js.Object must be embedded in a struct"
	_ ******js.Object       // want "js.Object must be embedded in a struct"
	_ []js.Object           // want "js.Object must be embedded in a struct"
	_ []*js.Object          // want "js.Object must be embedded in a struct"
	_ [10]js.Object         // want "js.Object must be embedded in a struct"
	_ [10]*js.Object        // want "js.Object must be embedded in a struct"
	_ map[string]js.Object  // want "js.Object must be embedded in a struct"
	_ map[string]*js.Object // want "js.Object must be embedded in a struct"
	_ map[js.Object]string  // want "js.Object must be embedded in a struct"
	_ map[*js.Object]string // want "js.Object must be embedded in a struct"
)

var (
	_ chan js.Object   // want "js.Object must be embedded in a struct"
	_ chan *js.Object  // want "js.Object must be embedded in a struct"
	_ chan []js.Object // want "js.Object must be embedded in a struct"

	_ = make(chan js.Object)           // want "js.Object must be embedded in a struct"
	_ = new(js.Object)                 // want "js.Object must be embedded in a struct"
	_ = new(*js.Object)                // want "js.Object must be embedded in a struct"
	_ = new(**js.Object)               // want "js.Object must be embedded in a struct"
	_ = make([]js.Object, 0, 10)       // want "js.Object must be embedded in a struct"
	_ = make(map[string]js.Object, 13) // want "js.Object must be embedded in a struct"
)
