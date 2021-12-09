package main

import (
	"net/http"

	"biplane.build"
	"biplane.build/server"
	"github.com/gorilla/mux"
)

// Create app used as main router
type App struct{}

// Routes method lets biplane use it as a router
func (a App) Routes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
}

// Take Off!
func main() {
	biplane.TakeOff(server.Config{
		Routers: []server.Router{
			new(App),
		},
	})
}
