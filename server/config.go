package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ccutch/biplane/database"
	"github.com/gorilla/mux"
)

// Main Config used for running server
type Config struct {
	Host string
	Port int

	Routers  []Router
	DBConfig database.Config
}

type Server struct {
	Config
}

func NewServer(h string, p int) *Server {
	return &Server{
		Config{Host: h, Port: p},
	}
}

func (s *Server) Controller(r Router) *Server {
	s.Routers = append(s.Routers, r)
	return s
}

func (s *Server) Database(c database.Config) *Server {
	s.DBConfig = c
	return s
}

func (s *Server) TakeOff() {
	r := mux.NewRouter()

	for _, o := range s.Routers {
		if c, ok := o.(Configurer); ok {
			c.Configure(s.Config)
		}

		o.Routes(r)
	}

	p := s.Port
	if p == 0 {
		p = 8080
	}

	u := fmt.Sprintf("%s:%d", s.Host, p)
	log.Printf("Server online and listening at %s", u)
	http.ListenAndServe(u, r)
}
