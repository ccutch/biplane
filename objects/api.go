package objects

import (
	"encoding/json"

	"github.com/ccutch/biplane/auth"
	"github.com/ccutch/biplane/database"
)

type Any = interface{}
type User = auth.User

type Parser func(Any) error
type Handler func(Parser) error

type API struct {
	Config database.Config
}

func NewClient(conf database.Config) API {
	return API{conf}
}

func stringify(d Any) string {
	b, err := json.Marshal(d)
	if err != nil {
		panic("Cannot serialize data")
	}

	return string(b)
}

func (a API) New(u User, k string, d Any) (*Object, error) {
	o := Object{
		Owner: u.ID,
		Kind:  k,
		Data:  stringify(d),
	}

	sql := `insert into objects (owner, kind, data)
					values ($1, $2, $3) returning id`

	return &o, a.Config.Client().
		QueryRow(sql, o.Owner, o.Kind, o.Data).
		Scan(&o.ID)
}

func (a API) Build(id int, u User, k string, d Any) *Object {
	return &Object{
		ID:    id,
		Owner: u.ID,
		Kind:  k,
		Data:  stringify(d),
	}
}

func (a API) Get(id int, d Any) (*Object, error) {
	o := Object{ID: id}
	sql := `select owner, kind, data from objects
					where id = $1`

	err := a.Config.Client().
		QueryRow(sql, o.ID).
		Scan(&o.Owner, &o.Kind, &o.Data)

	if err != nil {
		return nil, err
	}

	if d != nil {
		json.Unmarshal([]byte(o.Data), d)

		if c, ok := d.(ObjectReceiver); ok {
			c.SetObject(&o)
		}
	}

	return &o, nil
}

func (a API) ForUser(u User, k string, h Handler) ([]Object, error) {
	var objects []Object
	sql := `select id, data from objects
					where owner = $1 and kind = $2`

	r, err := a.Config.Client().Query(sql, u.ID, k)
	if err != nil {
		return objects, err
	}

	for r.Next() {
		o := Object{Owner: u.ID, Kind: k}
		err = r.Scan(&o.ID, &o.Data)

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

		objects = append(objects, o)
	}

	return objects, err

}
