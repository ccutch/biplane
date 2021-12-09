package main

import (
	"fmt"
	"net/http"

	"biplane.build"
	"biplane.build/mixins"
	"biplane.build/server"
)

type App struct {
	mixins.Controller
}

type Account struct {
	mixins.Model

	Name               string
	Avatar             string
	Age                int
	PrivateInformation string
}

type AccountLink struct {
	Link   string
	Avatar string
	Name   string
}

func (a Account) AccountLink() AccountLink {
	return AccountLink{
		Link:   fmt.Sprintf("/account?id=%d", a.ID()),
		Avatar: a.Avatar,
		Name:   a.Name,
	}
}

func (c App) Handler(w http.ResponseWriter, r *http.Request) {
	a := Account{
		Name:               "Connor",
		Avatar:             "data:jpg...",
		Age:                26,
		PrivateInformation: "This should not be given to a view!",
	}

	o, err := c.Objects().New(c.User(r), "Account", &a)
	fmt.Println("Object stored", o, err)

	c.Display(w, a.AccountLink())
}

func main() {
	biplane.TakeOff(server.Config{
		Routers: []server.Router{
			new(App),
		},
	})
}
