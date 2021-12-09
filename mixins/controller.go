package mixins

import (
	"net/http"

	"biplane.build/auth"
	"github.com/gorilla/mux"
)

type Controller struct {
	auth.Manager
}

func Routes(r *mux.Router)

func Handler(w http.ResponseWriter, r *http.Request)
