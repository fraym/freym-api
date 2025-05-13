package streams

import (
	"context"

	"github.com/fraym/freym-api/go/streams/domain/dto"
	"github.com/fraym/freym-api/go/streams/internal/client"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) GetAllEvents(
	ctx context.Context,
	tenant string,
	topic string,
	includedEventTypes []string,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	return c.Called(ctx, tenant, topic, includedEventTypes, perPage, queueSize, handler).Error(0)
}

func (c *MockClient) GetAllEventsAfterEvent(
	ctx context.Context,
	tenant string,
	topic string,
	includedEventTypes []string,
	eventId string,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	return c.Called(ctx, tenant, topic, includedEventTypes, eventId, perPage, queueSize, handler).Error(0)
}

func (c *MockClient) GetStream(
	ctx context.Context,
	tenant string,
	topic string,
	stream string,
	deploymentId int64,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	return c.Called(ctx, tenant, topic, stream, deploymentId, perPage, queueSize, handler).Error(0)
}

func (c *MockClient) GetStreamAfterEvent(
	ctx context.Context,
	tenant string,
	topic string,
	stream string,
	eventId string,
	deploymentId int64,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	return c.Called(ctx, tenant, topic, stream, eventId, deploymentId, perPage, queueSize, handler).Error(0)
}

func (c *MockClient) IsStreamEmpty(ctx context.Context, tenant string, topic string, stream string) (bool, error) {
	args := c.Called(ctx, tenant, topic, stream)
	return args.Bool(0), args.Error(1)
}

func (c *MockClient) GetEvent(
	ctx context.Context,
	tenantId string,
	topic string,
	eventId string,
) (*dto.SubscriptionEvent, error) {
	args := c.Called(ctx, tenantId, topic, eventId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.SubscriptionEvent), args.Error(1)
}

func (c *MockClient) GetLastEvent(ctx context.Context, tenantId string, topic string) (*dto.SubscriptionEvent, error) {
	args := c.Called(ctx, tenantId, topic)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.SubscriptionEvent), args.Error(1)
}

func (c *MockClient) GetLastEventByTypes(
	ctx context.Context,
	tenantId string,
	topic string,
	types []string,
) (*dto.SubscriptionEvent, error) {
	args := c.Called(ctx, tenantId, topic, types)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.SubscriptionEvent), args.Error(1)
}

func (c *MockClient) NewSubscription(
	topics []string,
	ignoreUnhandledEvents bool,
	deploymentId int64,
) *client.Subscription {
	args := c.Called(topics, ignoreUnhandledEvents, deploymentId)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*client.Subscription)
}

func (c *MockClient) Publish(ctx context.Context, topic string, events []*dto.PublishEvent) error {
	return c.Called(ctx, topic, events).Error(0)
}

func (c *MockClient) InvalidateGdprData(ctx context.Context, tenantId string, topic string, gdprId string) error {
	return c.Called(ctx, tenantId, topic, gdprId).Error(0)
}

func (c *MockClient) IntroduceGdprOnEventField(
	ctx context.Context,
	tenantId string,
	topic string,
	eventId string,
	fieldName string,
	defaultValue string,
) error {
	return c.Called(ctx, tenantId, topic, eventId, fieldName, defaultValue).Error(0)
}

func (c *MockClient) CreateStreamSnapshot(
	ctx context.Context,
	tenantId string,
	topic string,
	stream string,
	lastSnapshottedEventId string,
	snapshotEvent *dto.PublishEvent,
) error {
	return c.Called(ctx, tenantId, topic, stream, lastSnapshottedEventId, snapshotEvent).Error(0)
}

func (c *MockClient) RenameEventType(
	ctx context.Context,
	topic string,
	oldEventType string,
	newEventType string,
) error {
	return c.Called(ctx, topic, oldEventType, newEventType).Error(0)
}

func (c *MockClient) Close() {
	c.Called()
}
