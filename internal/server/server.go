package server

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"api/config"
	"api/internal/handler"
	"api/internal/router"
)

func Start(cfg *config.Config, userH *handler.UserHandler) error {
	srv := &http.Server{
		Addr:    cfg.HTTP.Port,
		Handler: router.Setup(userH),
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		slog.Info("shutting down...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			slog.Error("forced shutdown", "error", err)
		}
	}()

	slog.Info("server starting", "addr", cfg.HTTP.Port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
