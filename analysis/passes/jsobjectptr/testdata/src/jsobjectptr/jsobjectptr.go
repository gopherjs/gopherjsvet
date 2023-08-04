package directjsobject

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	_ js.Object                 // want "js.Object must always be a pointer"
	_ = js.Object{}             // want "js.Object must always be a pointer"
	_ []js.Object               // want "js.Object must always be a pointer"
	_ = []js.Object{}           // want "js.Object must always be a pointer"
	_ [10]js.Object             // want "js.Object must always be a pointer"
	_ = [10]js.Object{}         // want "js.Object must always be a pointer"
	_ map[int]js.Object         // want "js.Object must always be a pointer"
	_ = map[int]js.Object{}     // want "js.Object must always be a pointer"
	_ map[js.Object]int         // want "js.Object must always be a pointer"
	_ = map[js.Object]int{}     // want "js.Object must always be a pointer"
	_ chan js.Object            // want "js.Object must always be a pointer"
	_ = make(chan js.Object)    // want "js.Object must always be a pointer"
	_ struct{ js.Object }       // want "js.Object must always be a pointer"
	_ = struct{ js.Object }{}   // want "js.Object must always be a pointer"
	_ struct{ x js.Object }     // want "js.Object must always be a pointer"
	_ = struct{ x js.Object }{} // want "js.Object must always be a pointer"
	_ = new(js.Object)
	_ *js.Object
)

func _(js.Object) {} // want "js.Object must always be a pointer"

func _() js.Object { // want "js.Object must always be a pointer"
	return js.Object{} // want "js.Object must always be a pointer"
}

func _() {
	var x any
	_, _ = x.(js.Object) // want "js.Object must always be a pointer"
}

func _() {
	var x *js.Object
	_ = *x // TODO: This should fail
}
