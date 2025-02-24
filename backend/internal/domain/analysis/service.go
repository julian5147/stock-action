package analysis

import (
	"context"
	"sort"
	"stockapi/internal/domain/shared"
	"stockapi/internal/domain/stock"
	"time"
)

type AnalysisService struct {
	stockRepo stock.Repository
	logger    shared.Logger
}

func NewAnalysisService(repo stock.Repository, logger shared.Logger) *AnalysisService {
	return &AnalysisService{
		stockRepo: repo,
		logger:    logger,
	}
}

type StockAnalysis struct {
	Stock          *stock.Stock
	Score          float64
	Indicators     map[string]float64
	Recommendation string
}

func (s *AnalysisService) AnalyzeStocks(ctx context.Context) ([]StockAnalysis, error) {
	start := time.Now()
	stocks, err := s.stockRepo.FindAll(ctx)
	if err != nil {
		s.logger.Error(ctx, "Error fetching stocks for analysis", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}

	var analyses []StockAnalysis
	for _, stock := range stocks {
		analysis := s.analyzeStock(stock)
		analyses = append(analyses, analysis)
	}

	// Sort by score
	sort.Slice(analyses, func(i, j int) bool {
		return analyses[i].Score > analyses[j].Score
	})

	duration := time.Since(start)
	s.logger.Info(ctx, "Stock analysis completed", map[string]interface{}{
		"stocks_analyzed": len(stocks),
		"duration_ms":     duration.Milliseconds(),
	})
	return analyses, nil
}

func (s *AnalysisService) analyzeStock(stock *stock.Stock) StockAnalysis {
	score := stock.CalculateInvestmentScore()

	indicators := map[string]float64{
		"price_target_growth": calculatePriceTargetGrowth(stock),
		"rating_impact":       calculateRatingImpact(stock),
		"broker_confidence":   calculateBrokerConfidence(stock),
	}

	return StockAnalysis{
		Stock:          stock,
		Score:          score,
		Indicators:     indicators,
		Recommendation: determineRecommendation(score),
	}
}

func calculatePriceTargetGrowth(s *stock.Stock) float64 {
	return (s.Target.To.Amount - s.Target.From.Amount) / s.Target.From.Amount * 100
}

func calculateRatingImpact(s *stock.Stock) float64 {
	ratingScores := map[stock.Rating]float64{
		stock.Buy:     1.0,
		stock.Hold:    0.5,
		stock.Neutral: 0.3,
		stock.Sell:    0.0,
	}
	return ratingScores[s.Rating.To] - ratingScores[s.Rating.From]
}

func calculateBrokerConfidence(s *stock.Stock) float64 {
	// Implement logic based on the broker and its history
	return 0.75 // Example value
}

func determineRecommendation(score float64) string {
	switch {
	case score >= 0.8:
		return "Strong Buy"
	case score >= 0.6:
		return "Buy"
	case score >= 0.4:
		return "Hold"
	case score >= 0.2:
		return "Sell"
	default:
		return "Strong Sell"
	}
}
