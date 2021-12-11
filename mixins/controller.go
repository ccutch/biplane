package mixins

import (
	"encoding/json"
	"net/http"

	"github.com/ccutch/biplane/auth"
	"github.com/ccutch/biplane/objects"
	"github.com/ccutch/biplane/server"
	"github.com/gorilla/mux"
)

type Controller struct {
	auth.Manager

	config server.Config
}

func (c Controller) Routes(r *mux.Router) {
	r.HandleFunc("/", c.Handler)
}

func (c Controller) Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func (c Controller) Configure(conf server.Config) {
	c.config = conf
}

func (c Controller) Objects() objects.API {
	return objects.NewClient(c.config.DBConfig)
}

func (c Controller) Fail(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (c Controller) Display(w http.ResponseWriter, d interface{}) {
	err := json.NewEncoder(w).Encode(d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c Controller) ParseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
