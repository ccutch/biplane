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

type Greeting struct {
	mixins.View

	Greeting string
}

type App struct {
	mixins.Controller
}

func (a App) Handler(w http.ResponseWriter, r *http.Request) {
	var account Account
	var greeting Greeting

	_, err := a.Objects().Get(1, &account)
	if err != nil {
		a.Fail(err)
		return
	}

	greeting.Greeting = fmt.Sprintf("Hello %s", account.Name)
	greeting.Display(w)
}

func main() {
	biplane.TakeOff(server.Config{
		Port:    8080,
		Routers: []server.Router{new(App)},
	})
}
