package mixins

import "biplane.build/objects"

type Model struct {
	object *objects.Object
}

func SetObject(o *objects.Object)
