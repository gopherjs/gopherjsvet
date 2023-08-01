package directjsobject

import (
	"github.com/gopherjs/gopherjs/js"
)

type _ struct {
	Name       string `js:"name"`
	*js.Object        // want "js.Object must be first struct field"
}

type _ struct {
	Name       string
	*js.Object // This one is permitted, since there are no js tags in this struct
}

type _ struct {
	*js.Object        // Correct placement
	Name       string `js:"name"`
}
