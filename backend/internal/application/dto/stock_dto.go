package dto

import (
	"stockapi/internal/domain/analysis"
	"stockapi/internal/domain/stock"
	"time"
)

type StockResponse struct {
	ID         string    `json:"id"`
	Ticker     string    `json:"ticker"`
	TargetFrom float64   `json:"target_from"`
	TargetTo   float64   `json:"target_to"`
	Company    string    `json:"company"`
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	Time       time.Time `json:"time"`
}

type AnalysisResponse struct {
	Stock          StockResponse      `json:"stock"`
	Score          float64            `json:"score"`
	Indicators     map[string]float64 `json:"indicators"`
	Recommendation string             `json:"recommendation"`
}

func ToStockResponse(s *stock.Stock) StockResponse {
	return StockResponse{
		ID:         s.ID.String(),
		Ticker:     s.Ticker,
		TargetFrom: s.Target.From.Amount,
		TargetTo:   s.Target.To.Amount,
		Company:    s.Company,
		Action:     s.Action,
		Brokerage:  s.Brokerage,
		RatingFrom: string(s.Rating.From),
		RatingTo:   string(s.Rating.To),
		Time:       s.Time,
	}
}

func ToAnalysisResponse(analysis analysis.StockAnalysis) AnalysisResponse {
	return AnalysisResponse{
		Stock:          ToStockResponse(analysis.Stock),
		Score:          analysis.Score,
		Indicators:     analysis.Indicators,
		Recommendation: analysis.Recommendation,
	}
}
