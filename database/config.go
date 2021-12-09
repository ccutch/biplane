package database

import "database/sql"

type Config interface {
	Setup() error
	Client() *sql.DB
}
