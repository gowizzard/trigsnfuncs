// Copyright 2023 J&J Ideenschmiede GmbH. All rights reserved.

//go:build development

// Package environment_test is to test the development package.
package development_test

import (
	"encoding/json"
	"github.com/gowizzard/tfusion/files"
	"github.com/gowizzard/tfusion/internal/development"
	"strings"
	"syscall"
	"testing"
)

// TestEnvironment is to test the Import function.
func TestEnvironment(t *testing.T) {

	err := development.Environment()
	if err != nil {
		t.Error(err)
	}

	var data map[string]string
	err = json.Unmarshal(files.Development, &data)
	if err != nil {
		t.Error(err)
	}

	for key := range data {
		_, lookup := syscall.Getenv(strings.ToUpper(key))
		if !lookup {
			t.Errorf("the development variable %s was not found", key)
		}
	}

}
