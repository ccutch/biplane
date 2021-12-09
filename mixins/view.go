package mixins

import (
	"io"
)

type Viewable interface {
	Display(io.Writer)
}

type View struct{}

func (v View) Display(w io.Writer)
