package mixins

import (
	"encoding/json"
	"io"
	"net/http"

	"biplane.build/auth"
	"biplane.build/objects"
	"biplane.build/server"
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

func (c Controller) SetConfig(conf server.Config) {
	c.config = conf
}

func (c Controller) Objects() objects.API {
	return objects.NewClient(c.config.Database)
}

func (c Controller) Fail(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (c Controller) Display(w io.Writer, d Viewable) {
	d.Display(w)
}

func (c Controller) ParseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
