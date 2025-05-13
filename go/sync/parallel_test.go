package sync_test

import (
	"context"
	"errors"
	"sync"
	"testing"

	appSync "github.com/fraym/freym-api/go/sync"
	"github.com/fraym/freym-api/go/sync/internal/peer"
	"github.com/fraym/freym-api/go/sync/internal/peer/mocks"
	"github.com/fraym/freym-api/go/sync/publisher"
	"github.com/fraym/golog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type onSendToClientFn func(client peer.ServiceClient) (*publisher.SendResponse, error)

type mockBackchannelServiceClient struct {
	mock.Mock
}

func handleParallel(
	connections peer.PeerPool[peer.ServiceClient],
	onSendToClient onSendToClientFn,
) appSync.ParallelPublisher[peer.ServiceClient, any] {
	logger := golog.NewZerologLogger()

	sendFn := func(ctx context.Context, data any, client peer.ServiceClient) (*publisher.SendResponse, error) {
		return onSendToClient(client)
	}

	return appSync.NewParallelPublisher(logger, connections, sendFn)
}

func TestParallel(t *testing.T) {
	mockClient := &mockBackchannelServiceClient{}
	peerPool := &mocks.MockPool[peer.ServiceClient]{}
	peerPool.UseClient("client-1", nil)
	peerPool.UseClient("client-2", nil)
	peerPool.UseClient("client-3", mockClient)

	wg := sync.WaitGroup{}
	wg.Add(4)

	peerPool.On("Iterate", mock.Anything).Return(nil).Once()
	parallel := handleParallel(peerPool, func(peer.ServiceClient) (*publisher.SendResponse, error) {
		defer wg.Done()

		return &publisher.SendResponse{
			Ack: true,
		}, nil
	})

	err := parallel.Send(context.Background(), nil, func(address string) bool {
		return address != "client-2"
	}, func(_ string, response *publisher.Response) {
		defer wg.Done()
		assert.True(t, response.IsAcknowledged())
	})

	wg.Wait()
	assert.Nil(t, err)
	peerPool.AssertExpectations(t)
}

func TestParallel_NoClients(t *testing.T) {
	peerPool := &mocks.MockPool[peer.ServiceClient]{}

	peerPool.On("Iterate", mock.Anything).Return(nil).Once()
	parallel := handleParallel(peerPool, func(peer.ServiceClient) (*publisher.SendResponse, error) {
		assert.Fail(t, "not allowed to be called")

		return &publisher.SendResponse{}, nil
	})

	err := parallel.Send(context.Background(), nil, nil, nil)

	assert.Nil(t, err)
	peerPool.AssertExpectations(t)
}

func TestParallel_ClientLeavesTheCluster(t *testing.T) {
	mockClient := &mockBackchannelServiceClient{}
	peerPool := &mocks.MockPool[peer.ServiceClient]{}
	peerPool.UseClient("client-1", nil)
	peerPool.UseClient("client-2", mockClient)

	wasCalled := false
	wg := sync.WaitGroup{}
	wg.Add(3)

	peerPool.On("Iterate", mock.Anything).Return(nil).Once()
	peerPool.On("IsPeer", "client-2").Return(true)
	parallel := handleParallel(peerPool, func(client peer.ServiceClient) (*publisher.SendResponse, error) {
		defer wg.Done()

		if !wasCalled && client == mockClient {
			wasCalled = true
			return nil, errors.New("gone away")
		}

		return &publisher.SendResponse{
			Ack: true,
		}, nil
	})

	err := parallel.Send(context.Background(), nil, nil, nil)

	wg.Wait()
	assert.Nil(t, err)
	peerPool.AssertExpectations(t)
}

func TestParallel_ClientSendsNotAck(t *testing.T) {
	mockClient := &mockBackchannelServiceClient{}
	peerPool := &mocks.MockPool[peer.ServiceClient]{}
	peerPool.UseClient("client-1", nil)
	peerPool.UseClient("client-2", mockClient)

	wg := sync.WaitGroup{}
	wg.Add(2)

	peerPool.On("Iterate", mock.Anything).Return(nil).Once()
	parallel := handleParallel(peerPool, func(client peer.ServiceClient) (*publisher.SendResponse, error) {
		defer wg.Done()

		if client == mockClient {
			return &publisher.SendResponse{
				Ack: false,
			}, nil
		}

		return &publisher.SendResponse{
			Ack: true,
		}, nil
	})

	err := parallel.Send(context.Background(), nil, nil, nil)

	wg.Wait()
	assert.Nil(t, err)
	peerPool.AssertExpectations(t)
}
