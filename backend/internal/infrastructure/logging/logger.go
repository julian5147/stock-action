package logging

import (
	"context"
	"log"
	"stockapi/internal/domain/shared"
)

type StockLogger struct {
	logger *log.Logger
}

func NewStockLogger() shared.Logger {
	return &StockLogger{
		logger: log.Default(),
	}
}

func (l *StockLogger) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	l.Log(ctx, shared.INFO, msg, fields)
}

func (l *StockLogger) Log(ctx context.Context, level shared.LogLevel, msg string, fields map[string]interface{}) {
	l.logger.Printf("[%s] %s %v", level, msg, fields)
}

func (l *StockLogger) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	l.Log(ctx, shared.DEBUG, msg, fields)
}

func (l *StockLogger) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	l.Log(ctx, shared.ERROR, msg, fields)
}

func (l *StockLogger) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	l.Log(ctx, shared.WARNING, msg, fields)
}
