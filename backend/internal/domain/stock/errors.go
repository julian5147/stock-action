package stock

import "fmt"

type DomainError struct {
	Code    string
	Message string
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

var (
	// Data validation errors
	ErrInvalidTicker = &DomainError{
		Code:    "INVALID_TICKER",
		Message: "ticker cannot be empty or invalid",
	}

	ErrInvalidPrice = &DomainError{
		Code:    "INVALID_PRICE",
		Message: "price must be greater than zero",
	}

	ErrInvalidRating = &DomainError{
		Code:    "INVALID_RATING",
		Message: "rating must be one of: Buy, Sell, Hold, Neutral",
	}

	// Business rules errors
	ErrInvalidPriceTarget = &DomainError{
		Code:    "INVALID_PRICE_TARGET",
		Message: "target price must be different from current price",
	}

	ErrInvalidRatingTransition = &DomainError{
		Code:    "INVALID_RATING_TRANSITION",
		Message: "invalid rating transition detected",
	}

	ErrDuplicateAnalysis = &DomainError{
		Code:    "DUPLICATE_ANALYSIS",
		Message: "analysis already exists for this time period",
	}

	// Operation errors
	ErrStockNotFound = &DomainError{
		Code:    "STOCK_NOT_FOUND",
		Message: "stock not found in the system",
	}

	ErrAnalysisNotPossible = &DomainError{
		Code:    "ANALYSIS_NOT_POSSIBLE",
		Message: "insufficient data to perform analysis",
	}

	ErrInvalidTimeframe = &DomainError{
		Code:    "INVALID_TIMEFRAME",
		Message: "invalid timeframe for analysis",
	}

	// Limits and constraints errors
	ErrTooManyAnalysisRequests = &DomainError{
		Code:    "TOO_MANY_ANALYSIS_REQUESTS",
		Message: "exceeded maximum number of analysis requests",
	}

	ErrStaleData = &DomainError{
		Code:    "STALE_DATA",
		Message: "stock data is too old for accurate analysis",
	}
)
