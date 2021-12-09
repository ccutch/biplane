package server

import "github.com/gorilla/mux"

type Router interface {
	Routes(*mux.Router)
}

type Configurer interface {
	Configure(Config)
}
