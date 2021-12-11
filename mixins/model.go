package mixins

import (
	"github.com/ccutch/biplane/objects"
)

type Model struct {
	object *objects.Object
}

func (m *Model) SetObject(o *objects.Object) {
	m.object = o
}

func (m *Model) ID() int {
	return m.object.ID
}
