package database

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	client *sql.DB

	Host string
	Port int

	User     string
	Password string
	DBName   string
}

func (p Postgres) Setup() error {
	var err error

	p.client, err = sql.Open("postgres", p.connString())
	if err != nil {
		return err
	}

	_, err = p.client.Exec(createSQLTables())
	return err
}

func (p Postgres) Client() *sql.DB {
	if p.client == nil {
		panic("Postgres setup; no client available")
	}

	return p.client
}

func (p Postgres) connString() string {
	return fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=require",
		p.Host, p.Port, p.DBName, p.User, p.Password,
	)
}

func createSQLTables() string {
	return `
		create table if not exists users (
			id serial,
			username text unique,
			password text,
			enabled boolean,
			locked boolean
		);

		create table if not exists objects (
			id serial,
			owner int references users (id),
			kind text,
			data text
		);
	`
}
