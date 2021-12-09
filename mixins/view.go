package mixins

import "io"

type View struct {
}

func (v View) Display(w io.Writer)
