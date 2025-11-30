package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	api := application{
		config: cfg,
		db:     conn,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("Server failed starting", "error", err)
		os.Exit(1)
	}
}
