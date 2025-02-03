package subscription

import (
	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/golog"
)

type (
	OnSubscribedFn func(subscribed *managementpb.Subscribed) error
	OnPanicFn      func(panic *managementpb.Panic) error
	OnEventFn      func(event *managementpb.Event, connection Connection) error
)

type Endpoint struct {
	logger golog.Logger

	onSubscribedFn OnSubscribedFn
	onPanicFn      OnPanicFn
	onEventFn      OnEventFn
}

func NewEndpoint(logger golog.Logger) *Endpoint {
	return &Endpoint{
		logger:         logger,
		onSubscribedFn: defaultOnSubscribedFn,
		onPanicFn:      defaultOnPanicFn,
		onEventFn:      defaultOnEventFn,
	}
}

func (e *Endpoint) UseConnection(connection Connection) error {
	handler := e.handleRequest(connection)
	return connection.Handle(handler)
}

func (e *Endpoint) handleRequest(connection Connection) func(request *managementpb.SubscribeResponse) error {
	return func(response *managementpb.SubscribeResponse) error {
		if event := response.GetEvent(); event != nil {
			return e.onEventFn(event, connection)
		}

		if subscribed := response.GetSubscribed(); subscribed != nil {
			return e.onSubscribedFn(subscribed)
		}

		if panic := response.GetPanic(); panic != nil {
			return e.onPanicFn(panic)
		}

		return ErrUnknownGrpcRequest
	}
}

func (e *Endpoint) OnSubscribed(handler OnSubscribedFn) {
	if handler == nil {
		e.logger.Error().Write("recieved nil as OnSubscribedFn handler function for streams client endpoint")
		return
	}

	e.onSubscribedFn = handler
}

func (e *Endpoint) OnPanic(handler OnPanicFn) {
	if handler == nil {
		e.logger.Error().Write("recieved nil as OnEventFn handler function for streams client endpoint")
		return
	}

	e.onPanicFn = handler
}

func (e *Endpoint) OnEvent(handler OnEventFn) {
	if handler == nil {
		e.logger.Error().Write("recieved nil as OnEventFn handler function for streams client endpoint")
		return
	}

	e.onEventFn = handler
}
