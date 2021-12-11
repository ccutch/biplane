package objects

import "encoding/json"

type Object struct {
	api *API

	ID    int    `json:"id"`
	Owner int    `json:"owner"`
	Kind  string `json:"kind"`
	Data  string `json:"data"`
}

type ObjectReceiver interface {
	SetObject(o *Object)
}

func (o Object) Insert() error {
	sql := `insert into objects (id, owner, kind, data)
					values ($1, $2, $3, $4)`

	return o.api.Config.Client().
		QueryRow(sql, o.ID, o.Owner, o.Kind, o.Data).
		Err()
}

func (o Object) Refresh(d Any) error {
	sql := `select data objects where id = $1`
	return o.api.Config.Client().
		QueryRow(sql, o.ID).
		Scan(&o.Data)
}

func (o Object) Update(d Any) error {
	sql := `update objects set data = $2 where id = $1`
	o.Data = stringify(d)

	return o.api.Config.Client().QueryRow(sql, o.ID, o.Data).Err()
}

func (o Object) Delete() error {
	sql := `delete from objects where id = $1`
	_, err := o.api.Config.Client().Exec(sql, o.ID)
	return err
}

func (o Object) GetData(d Any) error {
	return json.Unmarshal([]byte(o.Data), d)
}
