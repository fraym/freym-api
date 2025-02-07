package peer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/fraym/freym-api/go/sync/config"
	"github.com/fraym/freym-api/go/sync/internal/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type ServiceClient any

// ConnectionHandlerFn returns true if the next connection can be handled.
// Return false if you do not want to handle the next connection.
type ConnectionHandlerFn[T ServiceClient] func(address string, client T) bool

type CreateServiceClientFn[T ServiceClient] func(cc grpc.ClientConnInterface) T

type PeerLeaveHandler func(address string)

type PeerPool[T ServiceClient] interface {
	IsPeer(address string) bool
	Iterate(ctx context.Context, handler ConnectionHandlerFn[T]) error
	OnPeerLeave(handler PeerLeaveHandler)
}

type pool[T ServiceClient] struct {
	mx                  *sync.RWMutex
	createServiceClient CreateServiceClientFn[T]
	config              *config.Config
	peerList            map[string]T
	peerConnList        map[string]*grpc.ClientConn
	service             *client.Service
	ownIp               string
}

func NewPool[T ServiceClient](
	ctx context.Context,
	conf *config.Config,
	service *client.Service,
	createServiceClient CreateServiceClientFn[T],
	ownIp string,
) PeerPool[T] {
	p := &pool[T]{
		mx:                  &sync.RWMutex{},
		createServiceClient: createServiceClient,
		config:              conf,
		peerList:            make(map[string]T),
		peerConnList:        make(map[string]*grpc.ClientConn),
		service:             service,
		ownIp:               ownIp,
	}

	_, _ = p.getPeer(p.ownIp)

	return p
}

func (p *pool[T]) IsPeer(address string) bool {
	return p.service.IsPeer(address)
}

func (p *pool[T]) Iterate(ctx context.Context, handler ConnectionHandlerFn[T]) error {
loop:
	for _, peerIp := range p.service.PeerIps() {
		select {
		case <-ctx.Done():
			break loop
		default:
			peer, err := p.getPeer(peerIp)
			if err != nil {
				return err
			}

			if !handler(peerIp, peer) {
				break loop
			}
		}
	}

	return nil
}

func (p *pool[T]) OnPeerLeave(handler PeerLeaveHandler) {
	p.service.OnPeerLeave(func(peerIp string) {
		p.mx.Lock()

		if peerConn, ok := p.peerConnList[peerIp]; ok {
			_ = peerConn.Close()
		}

		delete(p.peerList, peerIp)
		delete(p.peerConnList, peerIp)
		p.mx.Unlock()
	})
}

func (p *pool[T]) getPeer(address string) (T, error) {
	p.mx.RLock()
	peer, ok := p.peerList[address]
	p.mx.RUnlock()

	if ok {
		return peer, nil
	}

	p.mx.Lock()
	defer p.mx.Unlock()

	if client, ok := p.peerList[address]; ok {
		return client, nil
	}

	credentialsOptions := grpc.WithTransportCredentials(insecure.NewCredentials())
	keepaliveOptions := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                time.Minute,
		Timeout:             3 * time.Second,
		PermitWithoutStream: true,
	})

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", address, p.config.GrpcPort), credentialsOptions, keepaliveOptions)
	if err != nil {
		var peer T
		return peer, err
	}

	newPeerClient := p.createServiceClient(conn)

	p.peerConnList[address] = conn
	p.peerList[address] = newPeerClient

	return newPeerClient, nil
}
