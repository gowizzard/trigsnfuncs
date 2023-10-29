package database

import (
	"bytes"
	"errors"
)

// parse is to parse the serialization protocol from the redis database.
func parse(buf []byte) (s string, err error) {

	lines := bytes.Split(buf, []byte("\r\n"))
	for _, value := range lines {
		switch value[0] {
		case '+':
			return string(value[1:]), nil
		case '-':
			return "", errors.New(string(value[1:]))
		}
	}

	return "", nil

}
