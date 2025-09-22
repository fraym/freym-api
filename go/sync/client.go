package sync

import (
	"context"
	"net"

	"github.com/fraym/freym-api/go/sync/config"
	"github.com/fraym/freym-api/go/sync/internal/client"
	"github.com/fraym/freym-api/go/sync/internal/grpc"
	"github.com/fraym/freym-api/go/sync/internal/peer"
	"github.com/fraym/freym-api/go/sync/internal/util"
	"github.com/fraym/golog"
)

type Client[T peer.ServiceClient] interface {
	GetOwnAddress() string
	IteratePeers(ctx context.Context, handler peer.ConnectionHandlerFn[T]) error
	Lock(tenantId string, resource ...string) error
	TryLock(tenantId string, resource ...string) (bool, error)
	Unlock(tenantId string, resource ...string)
	RLock(tenantId string, resource ...string) error
	TryRLock(tenantId string, resource ...string) (bool, error)
	RUnlock(tenantId string, resource ...string)
	GetPeerPool() peer.PeerPool[T]
	Close() error
}

type syncClient[T peer.ServiceClient] struct {
	onStop   func() error
	ownIp    *net.IP
	service  *client.Service
	peerPool peer.PeerPool[T]
}

func NewClient[T peer.ServiceClient](
	logger golog.Logger,
	conf *config.Config,
	createServiceClient peer.CreateServiceClientFn[T],
) (Client[T], error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	ownIp, err := util.OwnIp()
	if err != nil {
		return nil, err
	}

	ctx, onStop := context.WithCancel(context.Background())

	serviceClient, stopServiceClient, err := grpc.NewClient(conf)
	if err != nil {
		return nil, handleStartupErr(logger, onStop, stopServiceClient, err)
	}

	service, err := client.NewService(ctx, serviceClient, logger, conf.AppPrefix, ownIp.String())
	if err != nil {
		return nil, handleStartupErr(logger, onStop, stopServiceClient, err)
	}

	return &syncClient[T]{
		onStop: func() error {
			onStop()
			service.WaitForStop()
			return stopServiceClient()
		},
		ownIp:    ownIp,
		service:  service,
		peerPool: peer.NewPool(ctx, conf, service, createServiceClient, ownIp.String()),
	}, nil
}

func (c *syncClient[T]) GetOwnAddress() string {
	return c.ownIp.String()
}

func (c *syncClient[T]) IteratePeers(ctx context.Context, handler peer.ConnectionHandlerFn[T]) error {
	return c.peerPool.Iterate(ctx, handler)
}

func (c *syncClient[T]) Lock(tenantId string, resource ...string) error {
	return c.service.Lock(tenantId, resource)
}

func (c *syncClient[T]) TryLock(tenantId string, resource ...string) (bool, error) {
	return c.service.TryLock(tenantId, resource)
}

func (c *syncClient[T]) Unlock(tenantId string, resource ...string) {
	c.service.Unlock(tenantId, resource)
}

func (c *syncClient[T]) RLock(tenantId string, resource ...string) error {
	return c.service.RLock(tenantId, resource)
}

func (c *syncClient[T]) TryRLock(tenantId string, resource ...string) (bool, error) {
	return c.service.TryRLock(tenantId, resource)
}

func (c *syncClient[T]) RUnlock(tenantId string, resource ...string) {
	c.service.RUnlock(tenantId, resource)
}

func (c *syncClient[T]) GetPeerPool() peer.PeerPool[T] {
	return c.peerPool
}

func (c *syncClient[T]) Close() error {
	return c.onStop()
}

func handleStartupErr(
	logger golog.Logger,
	onStop context.CancelFunc,
	stopServiceClient func() error,
	err error,
) error {
	onStop()

	if stopServiceClient != nil {
		if err2 := stopServiceClient(); err2 != nil {
			logger.Warn().WithError(err).Write("unable to close sync client")
		}
	}

	return err
}
