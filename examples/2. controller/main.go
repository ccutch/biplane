package main

import (
	"net/http"

	"biplane.build"
	"biplane.build/mixins"
	"biplane.build/server"
)

// This time we are going to use the controller mixin
// which will give us a Routes function
type App struct {
	mixins.Controller
}

// We are given a default handler with the controller mixin
// but we will define our own here
func (a App) Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My App Controller!"))
}

func main() {
	biplane.TakeOff(server.Config{
		Routers: []server.Router{
			new(App),
		},
	})
}
