package application

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type AppConfig struct {
	Addr string
	Mux  *http.ServeMux
}

type App struct {
	server *http.Server
}

func NewApplication(cfg *AppConfig) *App {

	return &App{server: &http.Server{Addr: ":" + cfg.Addr, Handler: cfg.Mux}}
}

func (app *App) Run(ctx context.Context, logger *slog.Logger, errCh chan<- error) {
	go func() {
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

}

func (app *App) Shutdown(ctx context.Context) error {
	shutCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	return app.server.Shutdown(shutCtx)
}
