package main

import (
	"log/slog"
	"os"

	"api/internal/model"
	"api/internal/server"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	app, err := InitializeApp()
	if err != nil {
		slog.Error("failed to initialize", "error", err)
		os.Exit(1)
	}

	if err := app.DB.AutoMigrate(&model.User{}); err != nil {
		slog.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}

	if err := server.Start(app.Config, app.UserHandler); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}
