package stock

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, stock *Stock) error
	FindByTicker(ctx context.Context, ticker string) (*Stock, error)
	FindAll(ctx context.Context) ([]*Stock, error)
	Update(ctx context.Context, stock *Stock) error
	Delete(ctx context.Context, id uuid.UUID) error
	Close(ctx context.Context) error
}

type StockAPIPort interface {
	FetchStocks(ctx context.Context) ([]*Stock, error)
}
