package client

import (
	"context"
	"sync"
	"time"

	"github.com/fraym/freym-api/go/proto/sync/managementpb"
	"github.com/fraym/freym-api/go/sync/internal/util"
	"github.com/fraym/golog"
	"github.com/samber/lo"
)

type lock struct {
	tenant   string
	resource []string
}

type lease struct {
	ctx         context.Context
	stopWg      *sync.WaitGroup
	mx          *sync.RWMutex
	cancelLease context.CancelFunc
	leaseId     string
	locks       []lock

	appPrefix string
	ttl       int32
	ownIp     string
	client    managementpb.ServiceClient
	logger    golog.Logger
	onRenew   func()
}

func NewLease(
	ctx context.Context,
	appPrefix string,
	ownIp string,
	client managementpb.ServiceClient,
	logger golog.Logger,
) *lease {
	stopWg := &sync.WaitGroup{}
	stopWg.Add(1)

	return &lease{
		ctx:         ctx,
		stopWg:      stopWg,
		mx:          &sync.RWMutex{},
		cancelLease: nil,
		leaseId:     "",
		locks:       nil,
		appPrefix:   appPrefix,
		ttl:         20,
		ownIp:       ownIp,
		client:      client,
		logger:      logger,
		onRenew:     nil,
	}
}

func (l *lease) OnRenew(fn func()) {
	l.mx.Lock()
	defer l.mx.Unlock()

	l.onRenew = fn
}

func (l *lease) LeaseId() string {
	l.mx.RLock()
	defer l.mx.RUnlock()

	return l.leaseId
}

func (l *lease) Disconnect() {
	l.mx.Lock()
	defer l.mx.Unlock()

	if l.cancelLease != nil {
		l.cancelLease()
		l.cancelLease = nil
	}
}

func (l *lease) WaitForStop() {
	l.stopWg.Wait()
}

func (l *lease) getRegisteredLocksForPb() []*managementpb.Lock {
	if len(l.locks) == 0 {
		return nil
	}

	return lo.Map(l.locks, func(lock lock, _ int) *managementpb.Lock {
		return managementpb.Lock_builder{
			LeaseId:  l.leaseId,
			TenantId: lock.tenant,
			Resource: lock.resource,
		}.Build()
	})
}

func (l *lease) Renew() error {
	l.mx.Lock()
	defer l.mx.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	l.cancelLease = cancel

	leaseResponse, err := l.client.CreateLease(l.ctx, managementpb.CreateLeaseRequest_builder{
		Ip:                     l.ownIp,
		App:                    l.appPrefix,
		Ttl:                    l.ttl,
		LeaseId:                l.leaseId,
		AlreadyRegisteredLocks: l.getRegisteredLocksForPb(),
	}.Build())
	if err != nil {
		return err
	}

	l.leaseId = leaseResponse.GetLeaseId()

	go l.keepLeaseAlive(ctx)

	l.logger.Debug().Write("renewed lease")
	return nil
}

func (l *lease) keepLeaseAlive(leaseCtx context.Context) {
	ticker := time.NewTicker(time.Duration(l.ttl/2) * time.Second)
	defer ticker.Stop()
	defer l.cancelLease()

	for {
		select {
		case <-leaseCtx.Done():
			return
		case <-l.ctx.Done():
			_, err := l.client.DropLease(context.Background(), managementpb.DropLeaseRequest_builder{
				LeaseId: l.LeaseId(),
			}.Build())
			if err != nil {
				l.logger.Warn().WithError(err).Write("unable to drop lease")
			}
			l.stopWg.Done()
			return
		case <-ticker.C:
			_, err := l.client.KeepLease(leaseCtx, managementpb.KeepLeaseRequest_builder{
				LeaseId: l.LeaseId(),
				Ttl:     l.ttl,
			}.Build())
			if err != nil {
				l.logger.Warn().WithError(err).Write("unable to keep lease")

				l.mx.RLock()
				l.onRenew()
				l.mx.RUnlock()

				if err := util.Retry(l.Renew, time.Millisecond*100, 50); err != nil {
					l.logger.Fatal().WithError(err).Write("unable to renew lease")
				}
				return
			}
		}
	}
}

func (l *lease) Track(tenant string, resource []string) {
	l.mx.Lock()
	defer l.mx.Unlock()

	l.locks = append(l.locks, lock{
		tenant:   tenant,
		resource: resource,
	})
}

func (l *lease) Untrack(tenant string, resource []string) {
	l.mx.Lock()
	defer l.mx.Unlock()

	l.locks = lo.Filter(l.locks, func(lock lock, _ int) bool {
		return lock.tenant != tenant || !util.SlicesEqual(lock.resource, resource)
	})
}
