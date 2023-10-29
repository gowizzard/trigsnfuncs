// Package database is to handle the redis database.
package database

// authenticate authenticates the client to the redis database.
func (c *Client) authenticate(username, password string) (response string, err error) {
	return c.execute("AUTH", username, password)
}
