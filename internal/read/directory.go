// Package read provides functions to read various types of files.
package read

import (
	"os"
	"path/filepath"
)

// Directory reads all JavaScript files from the specified directory and returns their contents as slices of bytes.
func Directory(name string) (scripts [][]byte, err error) {

	entry, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	for _, value := range entry {

		if !value.IsDir() && filepath.Ext(value.Name()) == ".js" {

			path := filepath.Join(name, value.Name())
			read, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}

			scripts = append(scripts, read)

		}

	}

	return scripts, nil

}
