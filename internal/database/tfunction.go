package database

func (c *Client) TFunctionLoad(s string) (response string, err error) {
	return c.execute("TFUNCTION", "LOAD", s)
}
