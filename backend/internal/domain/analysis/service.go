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
	LastUpdated    time.Time
}

// BrokerTier represents the prestige level of a broker
type BrokerTier int

const (
	TierS BrokerTier = iota + 1 // Global brokers of maximum prestige
	TierA                       // High prestige brokers
	TierB                       // Medium prestige brokers
	TierC                       // Low prestige brokers
)

// brokerTiers maps brokers to their prestige levels
var brokerTiers = map[string]BrokerTier{
	// Tier S - Global brokers of maximum prestige (1.0)
	"The Goldman Sachs Group": TierS,
	"Morgan Stanley":          TierS,
	"JPMorgan Chase & Co.":    TierS,
	"Bank of America":         TierS,
	"Citigroup":               TierS,

	// Tier A - High prestige brokers (0.8)
	"Wells Fargo & Company":            TierA,
	"UBS Group":                        TierA,
	"Deutsche Bank Aktiengesellschaft": TierA,
	"Barclays":                         TierA,
	"Royal Bank of Canada":             TierA,
	"HSBC":                             TierA,
	"BNP Paribas":                      TierA,
	"BMO Capital Markets":              TierA,
	"Mizuho":                           TierA,
	"Scotiabank":                       TierA,

	// Tier B - Established and specialized brokers (0.6)
	"Jefferies Financial Group": TierB,
	"Raymond James":             TierB,
	"Evercore ISI":              TierB,
	"Piper Sandler":             TierB,
	"TD Cowen":                  TierB,
	"Oppenheimer":               TierB,
	"Stifel Nicolaus":           TierB,
	"Keefe, Bruyette & Woods":   TierB,
	"Cantor Fitzgerald":         TierB,
	"Truist Financial":          TierB,
	"Wedbush":                   TierB,
	"Robert W. Baird":           TierB,
	"Sanford C. Bernstein":      TierB,
	"CIBC":                      TierB,
	"Macquarie":                 TierB,
	"Guggenheim":                TierB,
	"TD Securities":             TierB,
	"Susquehanna":               TierB,

	// Tier C - Boutique and regional brokers (0.4)
	"HC Wainwright":            TierC,
	"Stephens":                 TierC,
	"Roth Mkm":                 TierC,
	"Northland Securities":     TierC,
	"Benchmark":                TierC,
	"Chardan Capital":          TierC,
	"B. Riley":                 TierC,
	"Canaccord Genuity Group":  TierC,
	"Lake Street Capital":      TierC,
	"Leerink Partners":         TierC,
	"Loop Capital":             TierC,
	"DZ Bank":                  TierC,
	"KeyCorp":                  TierC,
	"DA Davidson":              TierC,
	"Lifesci Capital":          TierC,
	"BWS Financial":            TierC,
	"Wolfe Research":           TierC,
	"Rosenblatt Securities":    TierC,
	"Redburn Atlantic":         TierC,
	"Telsey Advisory Group":    TierC,
	"Craig Hallum":             TierC,
	"Maxim Group":              TierC,
	"JMP Securities":           TierC,
	"Argus":                    TierC,
	"Compass Point":            TierC,
	"LADENBURG THALM/SH SH":    TierC,
	"Tigress Financial":        TierC,
	"Alliance Global Partners": TierC,
	"Rodman & Renshaw":         TierC,
	"Fox Advisors":             TierC,
	"Glj Research":             TierC,
	"Westpark Capital":         TierC,
	"Hovde Group":              TierC,
	"Moffett Nathanson":        TierC,
	"Cfra":                     TierC,
	"CJS Securities":           TierC,
	"Northcoast Research":      TierC,
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

	if len(stocks) == 0 {
		return nil, stock.ErrAnalysisNotPossible
	}

	var analyses []StockAnalysis
	for _, stk := range stocks {
		// Check if data is not stale
		if time.Since(stk.Time) > 24*time.Hour {
			s.logger.Warn(ctx, "Stale data detected", map[string]interface{}{
				"stock_id":    stk.ID,
				"last_update": stk.Time,
			})
			continue
		}

		analysis, err := s.analyzeStock(ctx, stk)
		if err != nil {
			s.logger.Error(ctx, "Error analyzing stock", map[string]interface{}{
				"stock_id": stk.ID,
				"error":    err.Error(),
			})
			continue
		}
		analyses = append(analyses, analysis)
	}

	if len(analyses) == 0 {
		return nil, stock.ErrStaleData
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

func (s *AnalysisService) analyzeStock(ctx context.Context, stk *stock.Stock) (StockAnalysis, error) {
	// Validate that we have enough data for analysis
	if !s.hasRequiredData(stk) {
		return StockAnalysis{}, stock.ErrAnalysisNotPossible
	}

	score := stk.CalculateInvestmentScore()

	// Validate rating transition
	if !isValidRatingTransition(stk.Rating.From, stk.Rating.To) {
		return StockAnalysis{}, stock.ErrInvalidRatingTransition
	}

	// Validate price target
	if stk.Target.From.Amount == stk.Target.To.Amount {
		return StockAnalysis{}, stock.ErrInvalidPriceTarget
	}

	indicators := map[string]float64{
		"price_target_growth": calculatePriceTargetGrowth(stk),
		"rating_impact":       calculateRatingImpact(stk),
		"broker_confidence":   calculateBrokerConfidence(stk),
	}

	analysis := StockAnalysis{
		Stock:          stk,
		Score:          score,
		Indicators:     indicators,
		Recommendation: determineRecommendation(score),
		LastUpdated:    time.Now(),
	}

	return analysis, nil
}

func (s *AnalysisService) hasRequiredData(stk *stock.Stock) bool {
	return stk != nil &&
		stk.Target.From.Amount > 0 &&
		stk.Target.To.Amount > 0 &&
		stk.Rating.From != "" &&
		stk.Rating.To != ""
}

func isValidRatingTransition(from, to stock.Rating) bool {
	// No allow drastic changes from Buy to Sell without passing through Hold
	if from == stock.Buy && to == stock.Sell {
		return false
	}
	if from == stock.Sell && to == stock.Buy {
		return false
	}
	return true
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

func getPrestigeScore(brokerage string) float64 {
	tier, exists := brokerTiers[brokerage]
	if !exists {
		return 0.4 // Tier C por defecto
	}

	scores := map[BrokerTier]float64{
		TierS: 1.0,
		TierA: 0.8,
		TierB: 0.6,
		TierC: 0.4,
	}
	return scores[tier]
}

func getConsistencyScore(from, to stock.Rating) float64 {
	if from == to {
		return 1.0
	}
	if isValidRatingTransition(from, to) {
		return 0.7
	}
	return 0.3
}

func getPriceChangeScore(priceChange float64) float64 {
	switch {
	case priceChange >= -0.5 && priceChange <= 1.0:
		return 1.0
	case (priceChange > 1.0 && priceChange <= 2.0) || (priceChange < -0.5 && priceChange >= -0.7):
		return 0.7
	default:
		return 0.3
	}
}

func calculateBrokerConfidence(s *stock.Stock) float64 {
	prestigeScore := getPrestigeScore(s.Brokerage)
	consistencyScore := getConsistencyScore(s.Rating.From, s.Rating.To)
	priceChange := (s.Target.To.Amount - s.Target.From.Amount) / s.Target.From.Amount
	priceChangeScore := getPriceChangeScore(priceChange)

	finalScore := (prestigeScore * 0.4) + (consistencyScore * 0.3) + (priceChangeScore * 0.3)

	if finalScore > 1.0 {
		return 1.0
	}
	if finalScore < 0 {
		return 0
	}
	return finalScore
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

func (s *AnalysisService) AnalyzeAllStocks(ctx context.Context) ([]*stock.Stock, error) {
	stocks, err := s.stockRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// Sort by investment score
	sortedStocks := make([]*stock.Stock, len(stocks))
	copy(sortedStocks, stocks)

	for i := 0; i < len(sortedStocks)-1; i++ {
		for j := i + 1; j < len(sortedStocks); j++ {
			if sortedStocks[i].CalculateInvestmentScore() < sortedStocks[j].CalculateInvestmentScore() {
				sortedStocks[i], sortedStocks[j] = sortedStocks[j], sortedStocks[i]
			}
		}
	}

	return sortedStocks, nil
}
