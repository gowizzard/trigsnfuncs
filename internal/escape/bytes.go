// Package escape provides functionality to escape different types of characters.
package escape

import "strings"

// lookup table for escape characters.
var lookup = map[byte]string{
	'\\': `\\`,
	'"':  `\"`,
	'\n': `\n`,
	'\r': `\r`,
	'\t': `\t`,
}

// Bytes escapes the provided byte slice to a string, enclosing the result in double quotes.
func Bytes(v []byte) string {

	var build strings.Builder
	build.WriteByte('"')

	for _, value := range v {
		escape, ok := lookup[value]
		if ok {
			build.WriteString(escape)
		} else {
			build.WriteByte(value)
		}
	}

	build.WriteByte('"')
	return build.String()

}
