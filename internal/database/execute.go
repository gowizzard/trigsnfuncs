package database

import "strings"

// execute is to execute a command to the redis database.
func (c *Client) execute(s ...string) (response string, err error) {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	buf, err := c.command(strings.Join(s, " "))
	if err != nil {
		return "", err
	}

	return parse(buf)

}
