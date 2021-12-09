package prefab

import (
	"net/http"

	"biplane.build/mixins"
	"github.com/gorilla/mux"
)

type AuthController struct {
	mixins.Controller
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c AuthController) Routes(r *mux.Router) {
	s := r.NewRoute().Subrouter()

	s.Methods("POST").Path("/register").
		HandlerFunc(c.RegisterUser)

	s.Methods("POST").Path("/login").
		HandlerFunc(c.LoginUser)

	s.Methods("GET").Path("/tokens/validate").
		HandlerFunc(c.ValidateTokens)

	s.Methods("POST").Path("/tokens/exchange").
		HandlerFunc(c.ExchangeTokens)
}

func (c AuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("username")
	p := r.FormValue("password")

	user, err := c.Register(u, p)
	if err != nil {
		c.Fail(w, err)
		return
	}

	c.Display(w, AuthResponse{
		AccessToken:  user.AccessToken(),
		RefreshToken: user.RefreshToken(),
	})
}

func (c AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	user, err := c.Login(
		r.FormValue("username"),
		r.FormValue("password"),
	)

	if err != nil {
		c.Fail(w, err)
		return
	}

	c.Display(w, AuthResponse{
		AccessToken:  user.AccessToken(),
		RefreshToken: user.RefreshToken(),
	})
}

func (c AuthController) ValidateTokens(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("token")
	err := c.ValidateAccessToken(t)

	if err != nil {
		c.Fail(w, err)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func (c AuthController) ExchangeTokens(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("token")
	aT, rT, err := c.ConsumeRefreshToken(t)

	if err != nil {
		c.Fail(w, err)
		return
	}

	c.Display(w, AuthResponse{
		AccessToken:  aT,
		RefreshToken: rT,
	})
}
