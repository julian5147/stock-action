package application

import (
	"stockapi/internal/application/services"
	"stockapi/internal/domain/analysis"
)

type StockApplication struct {
	StockService    *services.StockService
	AnalysisService *analysis.AnalysisService
}

func NewStockApplication(
	stockService *services.StockService,
	analysisService *analysis.AnalysisService,
) *StockApplication {
	return &StockApplication{
		StockService:    stockService,
		AnalysisService: analysisService,
	}
}
