package catalog

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

// compile-time check
var _ CatalogRepository = (*Repository)(nil)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, p *Product) (*Product, error) {
	if p == nil {
		return nil, fmt.Errorf("product is nil")
	}

	// Adjust placeholders for your driver: "?"=MySQL, "$1"=Postgres, etc.
	res, err := r.db.ExecContext(ctx,
		"INSERT INTO products (name, stock_count, price_cents) VALUES ($1, $2, $3)",
		p.Name, p.StockCount, p.PriceCents)
	if err != nil {
		return nil, err
	}

	// Try to set ID from driver support; some drivers require RETURNING clause
	if id, err := res.LastInsertId(); err == nil {
		p.ID = id
	}

	return p, nil
}
