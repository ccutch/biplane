package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ccutch/biplane/database"
	"golang.org/x/crypto/bcrypt"
)

type Manager struct {
	Config database.Config
}

func NewManager(conf database.Config) *Manager {
	return &Manager{conf}
}

func (m *Manager) Register(u, p string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{
		Username: u,
		Password: string(hash),
		Enabled:  true,
	}

	sql := `insert into users (username, password)
					values ($1, $2) returning id`
	return &user, m.Config.Client().
		QueryRow(sql, user.Username, user.Password).
		Scan(&user.ID)
}

func (m *Manager) Login(u, p string) (*User, error) {
	user := User{Username: u}
	sql := `select password, enabled, locked from users
					where username = $1`

	err := m.Config.Client().
		QueryRow(sql, user.Username).
		Scan(&user.Password, &user.Enabled, &user.Locked)

	if err != nil {
		return nil, err
	}

	if user.Locked {
		return &user, fmt.Errorf("User locked")
	}

	if !user.Enabled {
		return &user, fmt.Errorf("User account disabled")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))
	if err != nil {
		return &user, fmt.Errorf("Invalid password")
	}

	return &user, nil
}

func (m *Manager) User(r *http.Request) User {
	if u, ok := r.Context().Value("user").(User); ok {
		return u
	}

	a := r.Header.Get("Authorization")
	parts := strings.Split(a, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		panic("Invalid `Authorization` header")
	}

	u, err := m.ParseAccessToken(parts[1])
	if err != nil {
		panic("Unable to parse token")
	}

	r.WithContext(context.WithValue(r.Context(), "user", *u))
	return *u
}
