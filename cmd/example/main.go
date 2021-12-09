package main

import (
	"biplane.build"
	"biplane.build/server"
)

func main() {
	biplane.TakeOff(server.Config{
		Port:    8080,
		Routers: []server.Router{},
	})
}
