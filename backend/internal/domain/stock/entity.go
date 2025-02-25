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
	// Very Positive Recommendations
	StrongBuy  Rating = "Strong-Buy"
	Outperform Rating = "Outperform"
	Overweight Rating = "Overweight"

	// Positive Recommendations
	Buy      Rating = "Buy"
	Positive Rating = "Positive"

	// Neutral Recommendations
	Hold          Rating = "Hold"
	Neutral       Rating = "Neutral"
	EqualWeight   Rating = "Equal-Weight"
	MarketPerform Rating = "Market-Perform"

	// Negative Recommendations
	Underweight  Rating = "Underweight"
	Underperform Rating = "Underperform"
	Sell         Rating = "Sell"
)

type Stock struct {
	ID        uuid.UUID
	Ticker    string
	Target    TargetPrice
	Company   string
	Action    string
	Brokerage string
	Rating    RatingChange
	Time      time.Time
}

type TargetPrice struct {
	From Money
	To   Money
}

type RatingChange struct {
	From Rating
	To   Rating
}

func NewStock(ticker string, targetFrom, targetTo Money) (*Stock, error) {
	if ticker == "" {
		return nil, errors.New("ticker cannot be empty")
	}

	now := time.Now()
	return &Stock{
		ID:     uuid.New(),
		Ticker: ticker,
		Target: TargetPrice{
			From: targetFrom,
			To:   targetTo,
		},
		Time: now,
	}, nil
}

func (s *Stock) CalculateInvestmentScore() float64 {
	var score float64

	// Factor 1: Growth Potential (30%)
	targetDiff := (s.Target.To.Amount - s.Target.From.Amount) / s.Target.From.Amount
	growthScore := targetDiff * 0.3

	// Factor 2: Broker Rating (25%)
	var ratingScore float64
	switch s.Rating.To {
	case StrongBuy, Outperform, Overweight:
		ratingScore = 0.25
	case Buy, Positive:
		ratingScore = 0.20
	case Hold, Neutral, EqualWeight, MarketPerform:
		ratingScore = 0.15
	case Underweight, Underperform:
		ratingScore = 0.05
	case Sell:
		ratingScore = 0
	}

	// Factor 3: Rating Improvement (15%)
	var ratingImprovementScore float64
	if s.Rating.From != s.Rating.To {
		if s.Rating.To == Buy && (s.Rating.From == Hold || s.Rating.From == Sell) {
			ratingImprovementScore = 0.15
		} else if s.Rating.To == Hold && s.Rating.From == Sell {
			ratingImprovementScore = 0.1
		}
	}

	// Factor 4: Broker Reputation (20%)
	var brokerageScore float64
	switch s.Brokerage {
	case "Goldman Sachs", "Morgan Stanley", "JP Morgan":
		brokerageScore = 0.2
	case "Bank of America", "Citigroup":
		brokerageScore = 0.15
	default:
		brokerageScore = 0.1
	}

	// Factor 5: Recommendation Timeliness (10%)
	var timelinessScore float64
	daysSinceUpdate := time.Since(s.Time).Hours() / 24
	if daysSinceUpdate <= 7 {
		timelinessScore = 0.1
	} else if daysSinceUpdate <= 30 {
		timelinessScore = 0.05
	}

	// Calculate final score
	score = growthScore + ratingScore + ratingImprovementScore + brokerageScore + timelinessScore

	// Normalize score between 0 and 1
	if score > 1 {
		score = 1
	} else if score < 0 {
		score = 0
	}

	return score
}
