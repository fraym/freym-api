package mocks

import (
	"context"

	"github.com/fraym/freym-api/go/sync/internal/peer"
	"github.com/stretchr/testify/mock"
)

type clientData[T peer.ServiceClient] struct {
	address string
	client  T
}

type MockPool[T peer.ServiceClient] struct {
	mock.Mock
	clients []clientData[T]
}

func (p *MockPool[T]) IsPeer(address string) bool {
	return p.Called(address).Bool(0)
}

func (p *MockPool[T]) Iterate(ctx context.Context, handler peer.ConnectionHandlerFn[T]) error {
loop:
	for _, clientData := range p.clients {
		select {
		case <-ctx.Done():
			break loop
		default:
			if !handler(clientData.address, clientData.client) {
				break loop
			}
		}
	}

	return p.Called(handler).Error(0)
}

func (p *MockPool[T]) OnPeerLeave(handler peer.PeerLeaveHandler) {
	p.Called(handler)
}

func (p *MockPool[T]) UseClient(address string, client T) {
	p.clients = append(p.clients, clientData[T]{
		address: address,
		client:  client,
	})
}
