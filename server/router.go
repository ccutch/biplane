package server

import "github.com/gorilla/mux"

type Router interface {
	Routes(*mux.Router)
}

type WithConfig interface {
	SetConfig(Config)
}
