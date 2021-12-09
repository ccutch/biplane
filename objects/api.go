package objects

import "biplane.build/database"

type API struct {
	Config database.Config
}

func NewClient(conf database.Config) API {
	return API{conf}
}

func (a API) New(u User, k string, d Any) (*Object, error)

func (a API) Build(id int, u User, k string, d Any) *Object

func (a API) Get(id int, d Any) (*Object, error)

func (a API) ForUser(u User, k string, h Handler) ([]Object, error)
