package directjsobject

import (
	"github.com/gopherjs/gopherjs/js"
)

type _ struct {
	Name       string `js:"name"`
	*js.Object        // want "*js.Object must be first struct field"
}
