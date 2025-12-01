package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed starting", "error", err)
		os.Exit(1)
	}
}
