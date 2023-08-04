package directjsobject

import (
	"github.com/gopherjs/gopherjs/js"
)

type _ struct {
	*js.Object
	Slice  []int       `js:"slice"`  // want "non-scalar types not permitted in structs that embed js.Object"
	Array  [10]int     `js:"array"`  // want "non-scalar types not permitted in structs that embed js.Object"
	Map    map[int]int `js:"map"`    // want "non-scalar types not permitted in structs that embed js.Object"
	Func   func()      `js:"func"`   // want "non-scalar types not permitted in structs that embed js.Object"
	Chan   chan int    `js:"chan"`   // want "non-scalar types not permitted in structs that embed js.Object"
	Struct struct{}    `js:"struct"` // want "non-scalar types not permitted in structs that embed js.Object"
}
