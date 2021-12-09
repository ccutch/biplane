package main

import (
	"net/http"

	"biplane.build"
	"biplane.build/mixins"
	"biplane.build/server"
)

//!!! Important Note !!!
// These are data views, meaning json not html.
// We dont support html yet, just use react imo

// We are going to use our mixin controller
// with a custom handler that displays our view
type App struct {
	mixins.Controller
}

// We make a structure representing data for our
// ui to consume. This would be for a React component
// that displays an avatar and a name and links to
// a users account
type AccountLink struct {
	mixins.View

	Link   string
	Avatar string
	Name   string
}

func (a App) Handler(w http.ResponseWriter, r *http.Request) {
	a.Display(w, AccountLink{
		Link:   "/my-profile",
		Avatar: "data:jpg...",
		Name:   "Connor",
	})
}

func main() {
	biplane.TakeOff(server.Config{
		Routers: []server.Router{
			new(App),
		},
	})
}
