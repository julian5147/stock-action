package services

import (
	"context"
	"fmt"
	"sort"
	"stockapi/internal/domain/shared"
	"stockapi/internal/domain/stock"
)

type StockService struct {
	repo    stock.Repository
	apiPort stock.StockAPIPort
	logger  shared.Logger
}

func NewStockService(repo stock.Repository, apiPort stock.StockAPIPort, logger shared.Logger) *StockService {
	return &StockService{
		repo:    repo,
		apiPort: apiPort,
		logger:  logger,
	}
}

func (s *StockService) SyncStocksFromAPI(ctx context.Context) error {
	s.logger.Info(ctx, "Starting stock synchronization from API", nil)

	stocks, err := s.apiPort.FetchStocks(ctx)
	if err != nil {
		s.logger.Error(ctx, "Failed to fetch stocks from API", map[string]interface{}{
			"error": err.Error(),
		})
		return fmt.Errorf("error fetching stocks: %w", err)
	}

	s.logger.Info(ctx, "Successfully fetched stocks from API", map[string]interface{}{
		"count": len(stocks),
	})

	for _, stock := range stocks {
		if err := s.repo.Save(ctx, stock); err != nil {
			s.logger.Error(ctx, "Failed to save stock", map[string]interface{}{
				"ticker": stock.Ticker,
				"error":  err.Error(),
			})
			return fmt.Errorf("error saving stock %s: %w", stock.Ticker, err)
		}
		s.logger.Debug(ctx, "Stock saved successfully", map[string]interface{}{
			"ticker": stock.Ticker,
			"id":     stock.ID,
		})
	}

	s.logger.Info(ctx, "Stock synchronization completed", map[string]interface{}{
		"total_synced": len(stocks),
	})
	return nil
}

func (s *StockService) GetRecommendedStocks(ctx context.Context) ([]*stock.Stock, error) {
	stocks, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching stocks: %w", err)
	}

	// Sort stocks by investment score
	sort.Slice(stocks, func(i, j int) bool {
		return stocks[i].CalculateInvestmentScore() > stocks[j].CalculateInvestmentScore()
	})

	return stocks, nil
}

func (s *StockService) GetAllStocks(ctx context.Context) ([]*stock.Stock, error) {
	stocks, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching all stocks: %w", err)
	}
	return stocks, nil
}
