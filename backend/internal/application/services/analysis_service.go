package services

import (
	"context"
	"stockapi/internal/application/dto"
	"stockapi/internal/application/events"
	"stockapi/internal/domain/analysis"
	"stockapi/internal/domain/stock"
)

type AnalysisApplicationService struct {
	analysisService *analysis.AnalysisService
	stockRepo       stock.Repository
	eventBus        events.EventBus
}

func NewAnalysisApplicationService(
	analysisService *analysis.AnalysisService,
	stockRepo stock.Repository,
	eventBus events.EventBus,
) *AnalysisApplicationService {
	return &AnalysisApplicationService{
		analysisService: analysisService,
		stockRepo:       stockRepo,
		eventBus:        eventBus,
	}
}

func (s *AnalysisApplicationService) AnalyzeAllStocks(ctx context.Context) ([]dto.AnalysisResponse, error) {
	analyses, err := s.analysisService.AnalyzeStocks(ctx)
	if err != nil {
		return nil, err
	}

	// Convert to DTOs and publish events
	var responses []dto.AnalysisResponse
	for _, analysis := range analyses {
		// Create and publish event
		event := stock.NewStockAnalyzedEvent(
			analysis.Stock.ID,
			analysis.Score,
			analysis.Indicators,
		)
		s.eventBus.Publish(ctx, "stock.analyzed", event)

		// Convert to DTO
		response := dto.ToAnalysisResponse(analysis)
		responses = append(responses, response)
	}

	return responses, nil
}
