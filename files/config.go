// Package files is to embed files to binary.
package files

import (
	_ "embed"
	"encoding/json"
)

//go:embed development/environment.json
var Development json.RawMessage
