package subscription

import (
	"context"
	"fmt"
	"sync"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/freym-api/go/streams/config"
	"github.com/fraym/freym-api/go/streams/domain/dto"
	"github.com/fraym/freym-api/go/streams/internal/util"
	"github.com/pkg/errors"
)

var (
	ErrSubscribeTimeout = newError("subscribe timeout")
	ErrClosed           = newError("closed")
	ErrNoEvent          = newError("no event")
)

type subscriber struct {
	sync.Mutex

	connection Connection
	endpoint   *Endpoint
	config     *config.Config

	subscribeLock    *sync.RWMutex
	subscribeErrChan chan string

	workCtx context.Context
	handler dto.HandlerFunc
	onPanic func(err error)
}

func NewSubscriber(
	ctx context.Context,
	connection Connection,
	endpoint *Endpoint,
	config *config.Config,
	parallelTopicProcessing bool,
	onPanic func(err error),
) *subscriber {
	s := &subscriber{
		connection:    connection,
		endpoint:      endpoint,
		config:        config,
		subscribeLock: &sync.RWMutex{},
		workCtx:       ctx,
		onPanic:       onPanic,
	}

	endpoint.OnSubscribed(s.handleSubscribed)
	endpoint.OnPanic(s.handlePanic)
	endpoint.OnEvent(s.handleEvent)

	return s
}

func (s *subscriber) handleEvent(
	subscribedEvent *managementpb.Event,
	connection Connection,
) error {
	if subscribedEvent == nil {
		return ErrNoEvent
	}

	options := subscribedEvent.GetOptions()

	if options != nil && options.GetBroadcast() {
		return s.handleBroadcastEvent(subscribedEvent)
	}

	return s.handleStreamedEvent(subscribedEvent, connection)
}

func (s *subscriber) handleBroadcastEvent(subscribedEvent *managementpb.Event) error {
	if s.handler == nil {
		return nil
	}

	subscriptionEvent, err := util.SubscriptionEventFromPb(subscribedEvent)
	if err != nil {
		return err
	}

	_, _ = s.handler(s.workCtx, subscriptionEvent)
	return nil
}

func (s *subscriber) handleStreamedEvent(subscribedEvent *managementpb.Event, connection Connection) error {
	if s.handler == nil {
		return connection.EventHandled(
			context.Background(),
			subscribedEvent.GetTenantId(),
			subscribedEvent.GetTopic(),
			"not subscribed",
			false,
			subscribedEvent.GetStream(),
		)
	}

	subscriptionEvent, err := util.SubscriptionEventFromPb(subscribedEvent)
	if err != nil {
		return err
	}

	if retry, err := s.handler(s.workCtx, subscriptionEvent); err != nil {
		return connection.EventHandled(
			context.Background(),
			subscribedEvent.GetTenantId(),
			subscribedEvent.GetTopic(),
			"error: "+err.Error(),
			retry,
			subscribedEvent.GetStream(),
		)
	}
	return connection.EventHandled(
		context.Background(),
		subscribedEvent.GetTenantId(),
		subscribedEvent.GetTopic(),
		"",
		false,
		subscribedEvent.GetStream(),
	)
}

func (s *subscriber) handleSubscribed(subscribed *managementpb.Subscribed) error {
	s.subscribeLock.RLock()
	defer s.subscribeLock.RUnlock()

	if s.subscribeErrChan != nil {
		s.subscribeErrChan <- subscribed.GetError()
	}

	return nil
}

func (s *subscriber) handlePanic(panicResponse *managementpb.Panic) error {
	s.onPanic(errors.New(panicResponse.GetReason()))
	return nil
}

func (s *subscriber) Subscribe(
	ctx context.Context,
	topics []string,
	parallelTopicProcessing bool,
	handler dto.HandlerFunc,
) error {
	groupId := s.config.GroupId
	if len(groupId) == 0 {
		return fmt.Errorf("cannot use empty value for STREAMS_CLIENT_GROUP_ID")
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, s.config.SendTimeout)
	defer cancel()

	s.subscribeLock.Lock()
	s.subscribeErrChan = make(chan string)
	s.subscribeLock.Unlock()

	if err := s.connection.Subscribe(ctx, groupId, topics, parallelTopicProcessing); err != nil {
		return err
	}

	select {
	case <-timeoutCtx.Done():
		return ErrSubscribeTimeout
	case <-s.workCtx.Done():
		return ErrClosed
	case subscribeErr := <-s.subscribeErrChan:
		if subscribeErr != "" {
			return errors.New(subscribeErr)
		}

		s.handler = handler
		return nil
	}
}
