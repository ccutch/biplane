package server

import (
	"biplane.build/database"
)

// Main Config used for running server
type Config struct {
	Host string
	Port int

	Routers  []Router
	Database database.Config
}

func NewServer(h string, p int, r ...Router) Config {
	return Config{
		Host: h,
		Port: p,

		Routers: r,
	}
}
