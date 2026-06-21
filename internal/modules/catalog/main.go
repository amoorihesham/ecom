package catalog

import (
	"database/sql"
	"log/slog"
	"net/http"
)

type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	StockCount int64  `json:"stock_count"`
	PriceCents int64  `json:"price_cents"`
}

func Initialize(mux *http.ServeMux, db *sql.DB, logger *slog.Logger) {
	repo := NewRepository(db)
	service := NewProductService(repo, logger)
	handlers := NewHandlers(service, logger)

	mux.HandleFunc("POST /catalog", handlers.Create)

}
