package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/fraym/freym-api/go/streams/domain/dto"
)

type handler struct {
	sync.RWMutex
	handler               map[string]dto.HandlerFunc
	handlerForAllTypes    dto.HandlerFunc
	ignoreUnhandledEvents bool
}

func NewHandler(ignoreUnhandledEvents bool) *handler {
	return &handler{
		handler:               map[string]dto.HandlerFunc{},
		ignoreUnhandledEvents: ignoreUnhandledEvents,
	}
}

func (h *handler) Use(eventType string, handlerFunc dto.HandlerFunc) {
	h.Lock()
	defer h.Unlock()

	h.handler[eventType] = handlerFunc
}

func (h *handler) UseForAllEventTypes(handlerFunc dto.HandlerFunc) {
	h.Lock()
	defer h.Unlock()

	h.handlerForAllTypes = handlerFunc
}

func (h *handler) Handle(ctx context.Context, event *dto.SubscriptionEvent) (bool, error) {
	h.RLock()
	defer h.RUnlock()

	select {
	case <-ctx.Done():
		return true, ctx.Err()
	default:
		if handler, ok := h.handler[event.Type]; ok {
			return handler(ctx, event)
		}

		if h.handlerForAllTypes != nil {
			return h.handlerForAllTypes(ctx, event)
		}

		if h.ignoreUnhandledEvents {
			return false, nil
		}

		return true, fmt.Errorf(
			"there is no event handler for the event of type '%s', topic '%s'",
			event.Type,
			event.Topic,
		)
	}
}
