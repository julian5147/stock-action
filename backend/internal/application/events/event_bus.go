package events

import (
	"context"
)

type EventBus interface {
	Publish(ctx context.Context, eventName string, event interface{}) error
	Subscribe(eventName string, handler func(ctx context.Context, event interface{}) error)
}
