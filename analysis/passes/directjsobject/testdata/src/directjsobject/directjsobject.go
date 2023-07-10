package directjsobject

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func directJSObject() {
	x := js.Object{} // want "js.Object must be embedded in a struct"
	fmt.Println(x)
}
