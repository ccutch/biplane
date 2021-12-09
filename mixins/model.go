package mixins

import (
	"encoding/json"
	"io"

	"biplane.build/objects"
)

type Model struct {
	object *objects.Object
}

func SetObject(o *objects.Object)

func (m *Model) ID() int {
	return m.object.ID
}

func (m *Model) Parse(r io.Reader) *Model {
	json.NewDecoder(r).Decode(m)
	return m
}
