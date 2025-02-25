package services

import (
	"context"
	"stockapi/internal/domain/analysis"
)

type AnalysisApplicationService struct {
	analysisService *analysis.AnalysisService
}

func NewAnalysisApplicationService(service *analysis.AnalysisService) *AnalysisApplicationService {
	return &AnalysisApplicationService{
		analysisService: service,
	}
}

func (s *AnalysisApplicationService) AnalyzeAllStocks(ctx context.Context) ([]analysis.StockAnalysis, error) {
	analyses, err := s.analysisService.AnalyzeStocks(ctx)
	if err != nil {
		return nil, err
	}
	return analyses, nil
}
