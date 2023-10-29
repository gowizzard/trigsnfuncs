package database

import (
	"context"
	"net"
	"sync"
)

// Options represents the options for the redis database.
type Options struct {
	Address  string
	Username string
	Password string
}

// Client represents the client for the redis database.
type Client struct {
	ctx   context.Context
	conn  net.Conn
	mutex sync.Mutex
}
