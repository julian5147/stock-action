package stock

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Money struct {
	Amount   float64
	Currency string
}

type Rating string

const (
	Buy     Rating = "Buy"
	Sell    Rating = "Sell"
	Hold    Rating = "Hold"
	Neutral Rating = "Neutral"
)

type Stock struct {
	ID     uuid.UUID
	Ticker string
	Target struct {
		From Money
		To   Money
	}
	Company   string
	Action    string
	Brokerage string
	Rating    struct {
		From Rating
		To   Rating
	}
	Time time.Time
}

func NewStock(ticker string, targetFrom, targetTo Money) (*Stock, error) {
	if ticker == "" {
		return nil, errors.New("ticker cannot be empty")
	}

	return &Stock{
		ID:     uuid.New(),
		Ticker: ticker,
		Target: struct {
			From Money
			To   Money
		}{
			From: targetFrom,
			To:   targetTo,
		},
		Time: time.Now(),
	}, nil
}

func (s *Stock) CalculateInvestmentScore() float64 {
	// Investment score calculation logic
	var score float64

	// Factor 1: Percentage difference between target
	targetDiff := (s.Target.To.Amount - s.Target.From.Amount) / s.Target.From.Amount
	score += targetDiff * 0.4 // 40% of the weight

	// Factor 2: Rating improvement
	if s.Rating.From != s.Rating.To {
		switch s.Rating.To {
		case Buy:
			score += 0.3
		case Hold:
			score += 0.1
		}
	}

	return score
}
