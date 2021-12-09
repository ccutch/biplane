package objects

// Edges are one -> many relations between objects
type Edge struct {
	Subject   Object
	Predicate string
	Object    Object
}

func (s Object) NewEdge(p string, o Object) (*Edge, error)

func (s Object) Edge(p string) ([]Object, error)

func (e Edge) Insert() error

func (e Edge) Delete() error
