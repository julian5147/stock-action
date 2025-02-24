package stock

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	GetID() uuid.UUID
	GetTime() time.Time
}

type StockAnalyzedEvent struct {
	ID         uuid.UUID
	StockID    uuid.UUID
	Score      float64
	Indicators map[string]float64
	Time       time.Time
}

func NewStockAnalyzedEvent(stockID uuid.UUID, score float64, indicators map[string]float64) *StockAnalyzedEvent {
	return &StockAnalyzedEvent{
		ID:         uuid.New(),
		StockID:    stockID,
		Score:      score,
		Indicators: indicators,
		Time:       time.Now(),
	}
}

func (e *StockAnalyzedEvent) GetID() uuid.UUID   { return e.ID }
func (e *StockAnalyzedEvent) GetTime() time.Time { return e.Time }
