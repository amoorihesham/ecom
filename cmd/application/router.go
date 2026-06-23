package application

import (
	"database/sql"
	"ecom/internal/modules/auth"
	"ecom/internal/shared/middlewares"
	"log/slog"
	"net/http"
)

type RouterConfig struct {
	Auth auth.AuthConfig
}

func NewRouter(db *sql.DB, logger *slog.Logger, cfg *RouterConfig) http.Handler {
	mux := http.NewServeMux()

	auth.Initialize(mux, db, logger, &cfg.Auth)

	wrappedMux := middlewares.Chain(middlewares.WithRequestId(), middlewares.WithLogging(logger))(mux)

	return wrappedMux

}
