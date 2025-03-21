package sync

import (
	"context"

	"github.com/fraym/freym-api/go/sync/internal/peer"
	"github.com/stretchr/testify/mock"
)

type MockClient[T peer.ServiceClient] struct {
	mock.Mock
}

func (c *MockClient[T]) GetOwnAddress() string {
	return c.Called().String(0)
}

func (c *MockClient[T]) IteratePeers(ctx context.Context, handler peer.ConnectionHandlerFn[T]) error {
	return c.Called(ctx, handler).Error(0)
}

func (c *MockClient[T]) Lock(tenantId string, resource ...string) error {
	return c.Called(tenantId, resource).Error(0)
}

func (c *MockClient[T]) Unlock(tenantId string, resource ...string) {
	c.Called(tenantId, resource)
}

func (c *MockClient[T]) RLock(tenantId string, resource ...string) error {
	return c.Called(tenantId, resource).Error(0)
}

func (c *MockClient[T]) RUnlock(tenantId string, resource ...string) {
	c.Called(tenantId, resource)
}

func (c *MockClient[T]) GetPeerPool() peer.PeerPool[T] {
	return c.Called().Get(0).(peer.PeerPool[T])
}

func (c *MockClient[T]) IsReady() bool {
	return c.Called().Bool(0)
}

func (c *MockClient[T]) Stop() error {
	return c.Called().Error(0)
}
