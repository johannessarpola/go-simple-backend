package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/johannessarpola/go-simple-backend/cmd"
)

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	slog.SetDefault(logger)

	cmd.Execute()

	for {
		time.Sleep(1000)
	}
}
