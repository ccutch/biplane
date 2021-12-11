package main

import (
	"net/http"

	"biplane.build"
	"biplane.build/mixins"
	"biplane.build/server"
	"github.com/gorilla/mux"
)

// This time we are going to use the controller mixin
// which will give us a Routes function
type App struct {
	mixins.Controller
}

// We are given a default handler with the controller mixin
// but we will define our own here
func (a App) Routes(r *mux.Router) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My App Controller!"))
	}

	r.Methods("GET").Path("/").HandlerFunc(handler)
}

func main() {
	biplane.TakeOff(server.Config{
		Routers: []server.Router{
			new(App),
		},
	})
}
