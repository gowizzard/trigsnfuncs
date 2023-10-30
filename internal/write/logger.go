// Package write is to write the log messages to the console.
package write

import (
	"log/slog"
	"os"
)

// options is to set the options for the handler.
// handler is to create a handler for the logger.
// Logger is to create a logger to handle the errors.
var (
	options = slog.HandlerOptions{
		AddSource: true,
	}
	handler = slog.NewTextHandler(os.Stdout, &options)
	Logger  = slog.New(handler)
)
