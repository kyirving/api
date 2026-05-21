package main

import (
	"log/slog"
	"os"

	"api/config"
	"api/internal/database"
	"api/internal/handler"
	"api/internal/repository"
	"api/internal/server"
	"api/internal/service"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	db, err := database.NewMySQL(cfg.MySQL)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	repo := repository.New(db)
	svc := service.New(repo)
	h := handler.New(svc)

	if err := server.Start(cfg, h); err != nil {
		slog.Error("server error", "error", err)
		os.Exit(1)
	}
}
