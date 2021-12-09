package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func (u User) AccessToken() string {
	c := jwt.StandardClaims{
		Subject:   fmt.Sprint(u.ID),
		ExpiresAt: time.Now().Unix() + int64(time.Hour*24*7),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	s, _ := t.SignedString(os.Getenv("JWT_KEY"))

	return s
}

func (u User) RefreshToken() string {
	c := jwt.StandardClaims{
		Subject:   fmt.Sprint(u.ID),
		ExpiresAt: time.Now().Unix() + int64(time.Hour*24*30),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	s, _ := t.SignedString(os.Getenv("JWT_KEY"))

	return s
}

func (m *Manager) getUser(id int) (*User, error) {
	user := User{ID: id}
	sql := `select username, password, enabled, locked from users
					where id = $1`

	return &user, m.Config.Client().
		QueryRow(sql, user.ID).
		Scan(&user.Username, &user.Password, &user.Enabled, &user.Locked)
}

func (m *Manager) ValidateAccessToken(t string) error {
	_, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_KEY"), nil
	})

	return err
}

func (m *Manager) ConsumeRefreshToken(t string) (string, string, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_KEY"), nil
	})

	if err != nil {
		return "", "", err
	}

	s := token.Claims.(jwt.StandardClaims).Subject
	id, err := strconv.Atoi(s)
	if err != nil {
		return "", "", err
	}

	u, err := m.getUser(id)
	if err != nil {
		return "", "", err
	}

	return u.AccessToken(), u.RefreshToken(), nil
}

func (m *Manager) ParseAccessToken(t string) (*User, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_KEY"), nil
	})

	if err != nil {
		return nil, err
	}

	s := token.Claims.(jwt.StandardClaims).Subject
	id, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	return m.getUser(id)
}
