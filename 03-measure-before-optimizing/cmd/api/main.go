package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thien/backend-learning-go/03-measure-before-optimizing/internal/app"
)

func main() {
	container := app.NewContainer()

	server := &http.Server{
		Addr:              ":8082",
		Handler:           container.Handler.Routes(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		container.Logger.Info("server starting", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			container.Logger.Error("server stopped unexpectedly", "error", err)
			os.Exit(1)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		container.Logger.Error("shutdown failed", "error", err)
		return
	}

	container.Logger.Info("server stopped cleanly")
	_ = slog.Default()
}
