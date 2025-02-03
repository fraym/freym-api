package streams

import (
	"context"

	"github.com/fraym/freym-api/go/streams/config"
	"github.com/fraym/freym-api/go/streams/domain/dto"
	"github.com/fraym/freym-api/go/streams/internal/client"
	"github.com/fraym/freym-api/go/streams/internal/grpc"
	"github.com/fraym/golog"
)

type Client interface {
	GetAllEvents(ctx context.Context, tenant string, topic string, includedEventTypes []string, perPage int, queueSize int, handler dto.HandlerFunc) error
	GetAllEventsAfterEvent(ctx context.Context, tenant string, topic string, includedEventTypes []string, eventId string, perPage int, queueSize int, handler dto.HandlerFunc) error
	GetStream(ctx context.Context, tenant string, topic string, stream string, perPage int, queueSize int, handler dto.HandlerFunc) error
	GetStreamAfterEvent(ctx context.Context, tenant string, topic string, stream string, eventId string, perPage int, queueSize int, handler dto.HandlerFunc) error
	IsStreamEmpty(ctx context.Context, tenant string, topic string, stream string) (bool, error)
	GetEvent(ctx context.Context, tenantId string, topic string, eventId string) (*dto.SubscriptionEvent, error)
	GetLastEvent(ctx context.Context, tenantId string, topic string) (*dto.SubscriptionEvent, error)
	GetLastEventByTypes(ctx context.Context, tenantId string, topic string, types []string) (*dto.SubscriptionEvent, error)
	NewSubscription(topics []string, ignoreUnhandledEvents bool) *client.Subscription
	Publish(ctx context.Context, topic string, events []*dto.PublishEvent) error
	InvalidateGdprData(ctx context.Context, tenantId string, topic string, gdprId string) error
	IntroduceGdprOnEventField(ctx context.Context, tenantId string, topic string, eventId string, fieldName string, defaultValue string) error
	CreateStreamSnapshot(ctx context.Context, tenantId string, topic string, stream string, lastSnapshottedEventId string, snapshotEvent *dto.PublishEvent) error
	Close()
}

type streamsClient struct {
	ctxCancel context.CancelFunc
	service   *client.Service
}

func NewClient(
	logger golog.Logger,
	conf *config.Config,
) (Client, error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	grpcClient, clientCloseFn, err := grpc.NewClient(conf)
	if err != nil {
		return nil, err
	}

	ctx, ctxCancel := context.WithCancel(context.Background())

	service, err := client.NewService(ctx, conf, grpcClient, logger)
	if err != nil {
		ctxCancel()
		if err2 := clientCloseFn(); err2 != nil {
			logger.Error().WithError(err).Write("unable to close streams client")
		}
		return nil, err
	}

	return &streamsClient{
		ctxCancel: ctxCancel,
		service:   service,
	}, nil
}

func (c *streamsClient) GetAllEvents(ctx context.Context, tenant string, topic string, includedEventTypes []string, perPage int, queueSize int, handler dto.HandlerFunc) error {
	return c.service.IterateAllEvents(ctx, tenant, topic, includedEventTypes, perPage, queueSize, handler)
}

func (c *streamsClient) GetAllEventsAfterEvent(ctx context.Context, tenant string, topic string, includedEventTypes []string, eventId string, perPage int, queueSize int, handler dto.HandlerFunc) error {
	return c.service.IterateAllEventsAfterEvent(ctx, tenant, topic, includedEventTypes, eventId, perPage, queueSize, handler)
}

func (c *streamsClient) GetStream(ctx context.Context, tenant string, topic string, stream string, perPage int, queueSize int, handler dto.HandlerFunc) error {
	return c.service.IterateStream(ctx, tenant, topic, stream, perPage, queueSize, handler)
}

func (c *streamsClient) GetStreamAfterEvent(ctx context.Context, tenant string, topic string, stream string, eventId string, perPage int, queueSize int, handler dto.HandlerFunc) error {
	return c.service.IterateStreamAfterEvent(ctx, tenant, topic, stream, eventId, perPage, queueSize, handler)
}

func (c *streamsClient) IsStreamEmpty(ctx context.Context, tenant string, topic string, stream string) (bool, error) {
	return c.service.IsStreamEmpty(ctx, tenant, topic, stream)
}

func (c *streamsClient) GetEvent(ctx context.Context, tenantId string, topic string, eventId string) (*dto.SubscriptionEvent, error) {
	return c.service.GetEvent(ctx, tenantId, topic, eventId)
}

func (c *streamsClient) GetLastEvent(ctx context.Context, tenantId string, topic string) (*dto.SubscriptionEvent, error) {
	return c.service.GetLastEvent(ctx, tenantId, topic)
}

func (c *streamsClient) GetLastEventByTypes(ctx context.Context, tenantId string, topic string, types []string) (*dto.SubscriptionEvent, error) {
	return c.service.GetLastEventByTypes(ctx, tenantId, topic, types)
}

func (c *streamsClient) NewSubscription(topics []string, ignoreUnhandledEvents bool) *client.Subscription {
	return c.service.NewSubscription(topics, ignoreUnhandledEvents)
}

func (c *streamsClient) Publish(ctx context.Context, topic string, events []*dto.PublishEvent) error {
	return c.service.Publish(ctx, topic, events)
}

func (c *streamsClient) InvalidateGdprData(ctx context.Context, tenantId string, topic string, gdprId string) error {
	return c.service.InvalidateGdprData(ctx, tenantId, topic, gdprId)
}

func (c *streamsClient) IntroduceGdprOnEventField(ctx context.Context, tenantId string, topic string, eventId string, fieldName string, defaultValue string) error {
	return c.service.IntroduceGdprOnEventField(ctx, tenantId, topic, eventId, fieldName, defaultValue)
}

func (c *streamsClient) CreateStreamSnapshot(ctx context.Context, tenantId string, topic string, stream string, lastSnapshottedEventId string, snapshotEvent *dto.PublishEvent) error {
	return c.service.CreateStreamSnapshot(ctx, tenantId, topic, stream, lastSnapshottedEventId, snapshotEvent)
}

func (c *streamsClient) Close() {
	c.ctxCancel()
}
