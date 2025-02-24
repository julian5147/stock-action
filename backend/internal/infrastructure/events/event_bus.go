package events

import (
	"context"
	"sync"
)

type Handler func(ctx context.Context, event interface{}) error

type EventBus struct {
	handlers map[string][]Handler
	mu       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]Handler),
	}
}

func (b *EventBus) Subscribe(eventType string, handler Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventType] = append(b.handlers[eventType], handler)
}

func (b *EventBus) Publish(ctx context.Context, eventType string, event interface{}) error {
	b.mu.RLock()
	handlers := b.handlers[eventType]
	b.mu.RUnlock()

	for _, handler := range handlers {
		if err := handler(ctx, event); err != nil {
			return err
		}
	}
	return nil
}
