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

)