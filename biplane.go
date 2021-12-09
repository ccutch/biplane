package biplane

import (
	"fmt"
	"log"
	"net/http"

	"biplane.build/server"
	"github.com/gorilla/mux"
)

// Start server with given server config
// This should also init the database client
// and setup any hooks.
func TakeOff(conf server.Config) {
	r := mux.NewRouter()
	u := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	for _, s := range conf.Routers {
		s.Routes(r)
	}

	log.Printf("Server online and listening at %s", u)
	http.ListenAndServe(u, r)
}
