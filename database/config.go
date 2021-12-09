package database

type Config interface {
	Setup() error
	ConnectionString() string
}
