package mixins

import (
	"encoding/json"
	"io"

	"biplane.build/objects"
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

func (m *Model) Parse(r io.Reader) *Model {
	json.NewDecoder(r).Decode(m)
	return m
}
