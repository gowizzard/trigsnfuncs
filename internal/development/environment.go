// Copyright 2023 J&J Ideenschmiede GmbH. All rights reserved.

//go:build development

// Package development is to import the development variables from the embedded data.
package development

import (
	"encoding/json"
	"github.com/gowizzard/tfusion/files"
	"strings"
	"syscall"
)

// Environment is to import the embedded development data into the development variables.
func Environment() error {

	var data map[string]string
	err := json.Unmarshal(files.Development, &data)
	if err != nil {
		return err
	}

	for key, value := range data {
		err = syscall.Setenv(strings.ToUpper(key), value)
		if err != nil {
			return err
		}
	}

	return nil

}
