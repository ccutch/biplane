package main

import (
	"fmt"
	"net/http"

	"biplane.build"
	"biplane.build/mixins"
	"biplane.build/server"
)

type Account struct {
	mixins.Model

	Name string
}

func (a Account) Greeting() Greeting {
	return Greeting{
		Message: fmt.Sprintf("Hello %s", a.Name),
	}
}

type Greeting struct {
	mixins.View

	Message string
}

type App struct {
	mixins.Controller
}

func (c App) Handler(w http.ResponseWriter, r *http.Request) {
	var a Account
	c.ParseJSON(r, &a)

	_, err := c.Objects().New(c.User(r), "Account", &a)
	if err != nil {
		c.Fail(w, err)
		return
	}

	c.Display(w, a.Greeting())
}

func main() {
	biplane.TakeOff(server.Config{
		Port:    8080,
		Routers: []server.Router{new(App)},
	})
}
