package objects

import "encoding/json"

// Properties are one -> one relations between objects
type Property struct {
	api *API

	Subject Object
	Name    string
	Object  Object
}

func (s Object) NewProperty(n string, o Object) (*Property, error) {
	p := Property{api: o.api, Subject: s, Name: n, Object: o}
	sql := `insert into properties (subject, name, object)
					value ($1, $2, $3)`

	return &p, s.api.Config.Client().
		QueryRow(sql, p.Subject.ID, p.Name, p.Object.ID).
		Scan()
}

func (s Object) Property(n string, h Handler) ([]Property, error) {
	var properties []Property
	sql := `
	select o.id, o.owner, o.kind, o.data
		from properties as e
		join objects as o on e.object = o.id
	where e.subject = $1 and e.name = $2`

	rows, err := s.api.Config.Client().Query(sql, s.ID, n)
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

		properties = append(properties, Property{s.api, s, n, o})
	}

	return properties, err

}

func (p Property) Update(o Object) error {
	p.Object = o
	sql := `update properties set object = $3
				  where subject = $1 and name = $2`

	return p.api.Config.Client().
		QueryRow(sql, p.Subject.ID, p.Name, p.Object.ID).
		Err()
}

func (p Property) Delete() error {
	sql := `delete from properties where subject = $1 and name = $2 and object = $3`
	return p.api.Config.Client().
		QueryRow(sql, p.Subject.ID, p.Name, p.Object.ID).
		Err()
}
