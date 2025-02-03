package client

import (
	"context"
	"sync"
	"time"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/freym-api/go/streams/config"
	"github.com/fraym/freym-api/go/streams/domain/dto"
	"github.com/fraym/freym-api/go/streams/internal/subscription"
	"github.com/fraym/freym-api/go/streams/internal/util"
	"github.com/fraym/golog"
	"github.com/pkg/errors"
)

type Subscription struct {
	mx        *sync.Mutex
	stop      context.CancelFunc
	parentCtx context.Context

	config  *config.Config
	client  managementpb.ServiceClient
	logger  golog.Logger
	handler *handler

	topics []string
}

func newSubscription(
	ctx context.Context,
	config *config.Config,
	client managementpb.ServiceClient,
	logger golog.Logger,
	topics []string,
	ignoreUnhandledEvents bool,
) *Subscription {
	return &Subscription{
		mx:        &sync.Mutex{},
		parentCtx: ctx,
		stop:      nil,
		config:    config,
		client:    client,
		logger:    logger,
		handler:   NewHandler(ignoreUnhandledEvents),
		topics:    topics,
	}
}

func (s *Subscription) UseHandler(eventType string, handlerFunc dto.HandlerFunc) {
	s.handler.Use(eventType, handlerFunc)
}

func (s *Subscription) UseHandlerForAllEventTypes(handlerFunc dto.HandlerFunc) {
	s.handler.UseForAllEventTypes(handlerFunc)
}

func (s *Subscription) Start() error {
	retryPause := time.Millisecond * 100
	subscriptionCtx, stop := context.WithCancel(s.parentCtx)
	defer stop()

	s.mx.Lock()
	s.stop = stop
	s.mx.Unlock()

	errChan := make(chan error)

	if err := util.Retry(func() error {
		return s.connect(subscriptionCtx, errChan)
	}, retryPause, 10); err != nil {
		return err
	}

	finalErrChan := make(chan error, 1)

	go func() {
		for {
			err := <-errChan

			if retryErr := util.Retry(func() error {
				return s.connect(subscriptionCtx, errChan)
			}, retryPause, 10); retryErr != nil {
				finalErrChan <- errors.Wrap(err, "unable to reconnect to streams service")
			}
		}
	}()

	select {
	case err := <-finalErrChan:
		return err
	case <-subscriptionCtx.Done():
		return nil
	}
}

func (s *Subscription) connect(subscriptionCtx context.Context, errChan chan error) error {
	connectionCtx, endConnection := context.WithCancel(subscriptionCtx)

	connection, err := subscription.NewConnection(connectionCtx, s.client, s.logger, s.config)
	if err != nil {
		defer endConnection()
		return err
	}

	endpoint := subscription.NewEndpoint(s.logger)
	subscriber := subscription.NewSubscriber(subscriptionCtx, connection, endpoint, s.config, func(err error) {
		defer endConnection()
		errChan <- err
	})

	go func() {
		if err := endpoint.UseConnection(connection); err != nil {
			defer endConnection()
			errChan <- err
		}
	}()

	if err := subscriber.Subscribe(connectionCtx, s.topics, s.handler.Handle); err != nil {
		endConnection()
		return err
	}

	return nil
}

func (s *Subscription) Stop() {
	s.mx.Lock()
	defer s.mx.Unlock()

	if s.stop != nil {
		s.stop()
	}
}
