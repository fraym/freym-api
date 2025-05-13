package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/fraym/freym-api/go/streams/domain/dto"
)

type handler struct {
	sync.RWMutex
	handlers              map[string][]dto.HandlerFunc
	handlersForAllTypes   []dto.HandlerFunc
	ignoreUnhandledEvents bool
}

func NewHandler(ignoreUnhandledEvents bool) *handler {
	return &handler{
		handlers:              make(map[string][]dto.HandlerFunc),
		ignoreUnhandledEvents: ignoreUnhandledEvents,
	}
}

func (h *handler) Use(eventType string, handlerFunc dto.HandlerFunc) {
	h.Lock()
	defer h.Unlock()

	h.handlers[eventType] = append(h.handlers[eventType], handlerFunc)
}

func (h *handler) UseForAllEventTypes(handlerFunc dto.HandlerFunc) {
	h.Lock()
	defer h.Unlock()

	h.handlersForAllTypes = append(h.handlersForAllTypes, handlerFunc)
}

func (h *handler) Handle(ctx context.Context, event *dto.SubscriptionEvent) error {
	h.RLock()
	defer h.RUnlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		var allHandlers []dto.HandlerFunc
		allHandlers = append(allHandlers, h.handlersForAllTypes...)

		handlers, ok := h.handlers[event.Type]
		if ok {
			allHandlers = append(allHandlers, handlers...)
		}
		if len(allHandlers) == 0 {
			if h.ignoreUnhandledEvents {
				return nil
			}

			return fmt.Errorf(
				"there is no event handler for the event of type '%s', topic '%s'",
				event.Type,
				event.Topic,
			)
		}

		wg := sync.WaitGroup{}
		var handlerErr error
		errLock := sync.Mutex{}

		for _, handler := range allHandlers {
			wg.Add(1)
			go func(handlerFunc dto.HandlerFunc) {
				defer wg.Done()
				if err := handlerFunc(ctx, event); err != nil {
					errLock.Lock()
					defer errLock.Unlock()
					handlerErr = err
				}
			}(handler)
		}

		wg.Wait()
		return handlerErr
	}
}
