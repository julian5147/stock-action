package shared

import (
	"context"
	"time"
)

type LogLevel string

const (
	DEBUG   LogLevel = "DEBUG"
	INFO    LogLevel = "INFO"
	WARNING LogLevel = "WARNING"
	ERROR   LogLevel = "ERROR"
)

type Logger interface {
	Log(ctx context.Context, level LogLevel, msg string, fields map[string]interface{})
	Debug(ctx context.Context, msg string, fields map[string]interface{})
	Info(ctx context.Context, msg string, fields map[string]interface{})
	Error(ctx context.Context, msg string, fields map[string]interface{})
	Warn(ctx context.Context, msg string, fields map[string]interface{})
}

type DomainLogger struct {
	logger Logger
}

func NewDomainLogger(logger Logger) *DomainLogger {
	return &DomainLogger{
		logger: logger,
	}
}

func (l *DomainLogger) LogStockAnalysis(ctx context.Context, stockID string, score float64, duration time.Duration) {
	l.logger.Info(ctx, "Stock analysis completed", map[string]interface{}{
		"stock_id": stockID,
		"score":    score,
		"duration": duration.String(),
	})
}

func (l *DomainLogger) LogError(ctx context.Context, err error, fields map[string]interface{}) {
	l.logger.Error(ctx, err.Error(), fields)
}

func (l *DomainLogger) Log(ctx context.Context, level LogLevel, msg string, fields map[string]interface{}) {
	l.logger.Log(ctx, level, msg, fields)
}

func (l *DomainLogger) Debug(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logger.Debug(ctx, msg, fields)
}

func (l *DomainLogger) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logger.Info(ctx, msg, fields)
}

func (l *DomainLogger) Warn(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logger.Warn(ctx, msg, fields)
}

func (l *DomainLogger) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logger.Error(ctx, msg, fields)
}
