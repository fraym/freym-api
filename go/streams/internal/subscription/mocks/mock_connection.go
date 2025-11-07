package mocks

import (
	"context"

	"github.com/fraym/freym-api/go/streams/internal/subscription"
	"github.com/stretchr/testify/mock"
)

type mockConnection struct {
	mock.Mock
}

func NewMockConnection() *mockConnection {
	return &mockConnection{}
}

func (c *mockConnection) Handle(handlerFn subscription.HandlerFn) error {
	c.Called(handlerFn)
	return nil
}

func (c *mockConnection) Close() {
	c.Called()
}

func (c *mockConnection) Subscribe(
	ctx context.Context,
	groupId string,
	topics []string,
	parallelTopicProcessing bool,
) error {
	return c.Called(ctx, groupId, topics, parallelTopicProcessing).Error(0)
}

func (c *mockConnection) EventHandled(
	ctx context.Context,
	tenantId string,
	topic string,
	errorMessage string,
	retry bool,
	stream string,
) error {
	return c.Called(ctx, tenantId, topic, errorMessage, retry, stream).Error(0)
}
