package application

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type AppConfig struct {
	Addr    string
	Handler http.Handler
}

type App struct {
	server *http.Server
}

func NewApplication(cfg *AppConfig) *App {

	return &App{server: &http.Server{Addr: ":" + cfg.Addr, Handler: cfg.Handler}}
}

func (app *App) Run(ctx context.Context, logger *slog.Logger, errCh chan<- error) {
	go func() {
		logger.Info("http server starting", "addr", app.server.Addr)
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		if err := app.server.Shutdown(shutdownCtx); err != nil {
			errCh <- err
		}
	}()

}

func (app *App) Shutdown(ctx context.Context) error {
	shutCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	return app.server.Shutdown(shutCtx)
}
