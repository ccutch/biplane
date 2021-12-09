package objects

// Properties are one -> one relations between objects
type Property struct {
	Subject Object
	Name    string
	Object  Object
}

func (s Object) NewProperty(p string, o Object) (*Property, error)

func (s Object) Property(p string) (*Object, error)

func (p Property) Insert() error

func (p Property) Update(o Object) error

func (p Property) Delete() error
