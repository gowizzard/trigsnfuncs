// Package main is the entrypoint for the hole project.
package main

import (
	"context"
	"github.com/gowizzard/trigsnfuncs/internal/database"
	"github.com/gowizzard/trigsnfuncs/internal/escape"
	"github.com/gowizzard/trigsnfuncs/internal/read"
	"github.com/gowizzard/trigsnfuncs/internal/write"
	"os"
)

// path is to set the path to the directory where the redis triggers or functions are stored.
const (
	path = "/scripts"
)

// address, username and password are to set the environment variables.
var (
	address  = os.Getenv("REDIS_ADDRESS")
	username = os.Getenv("REDIS_USERNAME")
	password = os.Getenv("REDIS_PASSWORD")
)

// init is to check if the environment variables are set.
func init() {

	if address == "" {
		write.Logger.Error("Environment variable REDIS_ADDRESS is not set.")
		os.Exit(1)
	}

}

// main is to load the redis triggers or functions to the database.
func main() {

	c, err := database.NewClient(context.Background(), database.Options{
		Address:  address,
		Username: username,
		Password: password,
	})
	if err != nil {
		write.Logger.Error("Create redis database client.", err)
		os.Exit(1)
	}

	defer func(c *database.Client) {
		err := c.Close()
		if err != nil {
			write.Logger.Error("Close redis database connection.", err)
			os.Exit(1)
		}
	}(c)

	scripts, err := read.Directory(path)
	if err != nil {
		write.Logger.Error("Read redis triggers or functions from directory.", err)
		os.Exit(1)
	}

	for _, value := range scripts {

		_, err = c.TFunctionLoad(escape.Bytes(value))
		if err != nil {
			write.Logger.Error("Load redis trigger or function to database.", err)
		}

	}

}
