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
		// Define the rating levels to calculate the improvement
		ratingLevels := map[Rating]int{
			// Very Positive (level 4)
			StrongBuy:  4,
			Outperform: 4,
			Overweight: 4,

			// Positive (level 3)
			Buy:      3,
			Positive: 3,

			// Neutral (level 2)
			Hold:          2,
			Neutral:       2,
			EqualWeight:   2,
			MarketPerform: 2,

			// Negative (level 1)
			Underweight:  1,
			Underperform: 1,
			Sell:         1,
		}

		fromLevel := ratingLevels[s.Rating.From]
		toLevel := ratingLevels[s.Rating.To]

		// Calculate the improvement based on the levels
		levelImprovement := toLevel - fromLevel

		switch {
		case levelImprovement >= 3:
			ratingImprovementScore = 0.15 // Maximum improvement (e.g., from Sell to StrongBuy)
		case levelImprovement == 2:
			ratingImprovementScore = 0.12 // Significant improvement (e.g., from Sell to Buy)
		case levelImprovement == 1:
			ratingImprovementScore = 0.08 // Moderate improvement (e.g., from Sell to Hold)
		case levelImprovement < 0:
			ratingImprovementScore = 0.0 // No improvement, it's a downgrade
		}
	}

	// Factor 4: Broker Reputation (20%)
	var brokerageScore float64
	switch s.Brokerage {
	// Tier S - Global brokers of maximum prestige
	case "The Goldman Sachs Group", "Morgan Stanley", "JPMorgan Chase & Co.", "Bank of America", "Citigroup":
		brokerageScore = 0.20

	// Tier A - High prestige brokers
	case "Wells Fargo & Company", "UBS Group", "Deutsche Bank Aktiengesellschaft",
		"Barclays", "Royal Bank of Canada", "HSBC", "BNP Paribas",
		"BMO Capital Markets", "Mizuho", "Scotiabank":
		brokerageScore = 0.15

	// Tier B - Established and specialized brokers
	case "Jefferies Financial Group", "Raymond James", "Evercore ISI",
		"Piper Sandler", "TD Cowen", "Oppenheimer", "Stifel Nicolaus",
		"Keefe, Bruyette & Woods", "Cantor Fitzgerald", "Truist Financial",
		"Wedbush", "Robert W. Baird", "Sanford C. Bernstein", "CIBC",
		"Macquarie", "Guggenheim", "TD Securities", "Susquehanna":
		brokerageScore = 0.10

	// Tier C - Boutique and regional brokers (y cualquier otro broker no listado)
	default:
		brokerageScore = 0.05
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
