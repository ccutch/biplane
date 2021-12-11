package objects

import "encoding/json"

// Edges are one -> many relations between objects
type Edge struct {
	api *API

	Subject   Object
	Predicate string
	Object    Object
}

func (s Object) NewEdge(p string, o Object) (*Edge, error) {
	e := Edge{api: o.api, Subject: s, Predicate: p, Object: o}
	sql := `insert into edges (subject, predicate, object)
					value ($1, $2, $3)`

	return &e, s.api.Config.Client().
		QueryRow(sql, e.Subject.ID, e.Predicate, e.Object.ID).
		Scan()
}

func (s Object) Edge(p string, h Handler) ([]Edge, error) {

	var edges []Edge
	sql := `
	select o.id, o.owner, o.kind, o.data
		from edges as e
		join objects as o on e.object = o.id
	where e.subject = $1 and e.predicate = $2`

	rows, err := s.api.Config.Client().Query(sql, s.ID, p)
	for rows.Next() {
		o := Object{api: s.api}

		err = rows.Scan(&o.ID, &o.Owner, &o.Kind, &o.Data)
		if err != nil {
			break
		}

		err = h(func(a Any) error {
			err := json.Unmarshal([]byte(o.Data), a)

			if err != nil {
				return err
			}

			if c, ok := a.(ObjectReceiver); ok {
				c.SetObject(&o)
			}

			return nil
		})

		if err != nil {
			break
		}

		edges = append(edges, Edge{s.api, s, p, o})
	}

	return edges, err
}

func (e Edge) Delete() error {
	sql := `delete from edges where subject = $1 and predicate = $2 and object = $3`
	_, err := e.api.Config.Client().Exec(sql, e.Subject.ID, e.Predicate, e.Object.ID)
	return err
}
