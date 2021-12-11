package biplane

import (
	"github.com/ccutch/biplane/server"
)

// Start server with given server config
// This should also init the database client
// and setup any hooks.
func NewServer(h string, p int) *server.Server {
	return server.NewServer(h, p)
}
