package database

// Close closes the connection to the redis database.
func (c *Client) Close() error {

	err := c.conn.Close()
	if err != nil {
		return err
	}

	return nil

}
