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
)
