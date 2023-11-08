package database

import (
	"context"
	"net"
)

// NewClient creates a new client for the redis database with user authentication.
func NewClient(ctx context.Context, o Options) (client *Client, err error) {

	conn, err := (&net.Dialer{}).DialContext(ctx, "tcp", o.Address)
	if err != nil {
		return nil, err
	}

	if o.Username != "" || o.Password != "" {
		_, err = (&Client{ctx: ctx, conn: conn}).authenticate(o.Username, o.Password)
		if err != nil {
			return nil, err
		}
	}

	return &Client{ctx: ctx, conn: conn}, nil

}
