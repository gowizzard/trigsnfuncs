package database

// Ping is to ping the redis database.
func (c *Client) Ping() (response string, err error) {
	return c.execute("PING")
}
