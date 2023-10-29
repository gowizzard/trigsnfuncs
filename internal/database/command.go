package database

// command sends a command to the redis server and returns the response.
func (c *Client) command(s string) (buf []byte, err error) {

	err = c.ctx.Err()
	if err != nil {
		return nil, err
	}

	_, err = c.conn.Write([]byte(s + "\r\n"))
	if err != nil {
		return nil, err
	}

	buf = make([]byte, 1024)
	n, err := c.conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil

}
