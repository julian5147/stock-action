package application

import (
	"stockapi/internal/application/services"
	"stockapi/internal/domain/analysis"
	"stockapi/internal/domain/shared"
	"stockapi/internal/domain/stock"
)

type StockApplication struct {
	StockService    *services.StockService
	AnalysisService *services.AnalysisApplicationService
}

func NewStockApplication(
	stockRepo stock.Repository,
	stockAPI stock.StockAPIPort,
	logger shared.Logger,
) *StockApplication {
	analysisService := analysis.NewAnalysisService(stockRepo, logger)
	return &StockApplication{
		StockService:    services.NewStockService(stockRepo, stockAPI, logger),
		AnalysisService: services.NewAnalysisApplicationService(analysisService),
	}
}
