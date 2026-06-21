package catalog

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type CatalogRepository interface {
	Create(ctx context.Context, p *Product) (*Product, error)
}

type CatalogService interface {
	Create(ctx context.Context, p *Product) (*Product, error)
}

type ProductService struct {
	repo   CatalogRepository
	logger *slog.Logger
}

// Compile-time assertion: ProductService implements CatalogService
var _ CatalogService = (*ProductService)(nil)

func NewProductService(repo CatalogRepository, logger *slog.Logger) *ProductService {
	return &ProductService{repo: repo, logger: logger}
}

func (s *ProductService) Create(ctx context.Context, p *Product) (*Product, error) {
	if p == nil {
		return nil, fmt.Errorf("product is nil")
	}
	if p.Name == "" {
		return nil, fmt.Errorf("product name is required")
	}

	created, err := s.repo.Create(ctx, p)
	if err != nil {
		if s.logger != nil {
			s.logger.Error("repository Create failed", "err", err)
		}
		return nil, err
	}
	time.Sleep(10 * time.Second)
	return created, nil
}
