//go:build development

// Package database_test is to test the database package.
package database_test

import (
	"context"
	"github.com/gowizzard/trigsnfuncs/internal/database"
	"github.com/gowizzard/trigsnfuncs/internal/development"
	"os"
	"testing"
)

// redis is to save the client for the redis database.
var (
	c *database.Client
)

// TestMain is the entrypoint for the tests.
func TestMain(m *testing.M) {

	err := development.Environment()
	if err != nil {
		panic(err)
	}

	c, err = database.NewClient(context.Background(), database.Options{
		Address:  os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	if err != nil {
		panic(err)
	}

	m.Run()

	err = c.Close()
	if err != nil {
		panic(err)
	}

}
