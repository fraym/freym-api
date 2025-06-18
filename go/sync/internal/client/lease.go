package client

import (
	"context"
	"sync"
	"time"

	"github.com/fraym/freym-api/go/proto/sync/managementpb"
	"github.com/fraym/freym-api/go/sync/internal/util"
	"github.com/fraym/golog"
)

type lock struct {
	lockId     string
	ctx        context.Context
	stopWg     *sync.WaitGroup
	mx         *sync.RWMutex
	cancelLock context.CancelFunc
	tenant     string
	resource   []string

	appPrefix string
	ttl       int32
	ownIp     string
	client    managementpb.ServiceClient
	logger    golog.Logger
	onRenew   func()
}

// type lease struct {
// 	ctx         context.Context
// 	stopWg      *sync.WaitGroup
// 	mx          *sync.RWMutex
// 	cancelLease context.CancelFunc
// 	leaseId     string
// 	locks       []lock
// 	rLocks      []lock

// 	appPrefix string
// 	ttl       int32
// 	ownIp     string
// 	client    managementpb.ServiceClient
// 	logger    golog.Logger
// 	onRenew   func()
// }

func NewLock(
	ctx context.Context,
	appPrefix string,
	ownIp string,
	client managementpb.ServiceClient,
	logger golog.Logger,
) *lock {
	stopWg := &sync.WaitGroup{}
	stopWg.Add(1)

	return &lock{
		lockId:     "",
		ctx:        ctx,
		stopWg:     stopWg,
		mx:         &sync.RWMutex{},
		cancelLock: nil,
		// locks:      nil,
		// rLocks:     nil,
		appPrefix: appPrefix,
		ttl:       20,
		ownIp:     ownIp,
		client:    client,
		logger:    logger,
		onRenew:   nil,
	}
}

func (l *lock) OnRenew(fn func()) {
	l.mx.Lock()
	defer l.mx.Unlock()

	l.onRenew = fn
}

func (l *lock) LockId() string {
	l.mx.RLock()
	defer l.mx.RUnlock()

	return l.lockId
}

func (l *lock) WaitForStop() {
	l.stopWg.Wait()
}

// func (l *lock) getRegisteredLocksForPb() []*managementpb.Lock {
// 	if len(l.locks) == 0 {
// 		return nil
// 	}

// 	return lo.Map(l.locks, func(lock lock, _ int) *managementpb.Lock {
// 		return managementpb.Lock_builder{
// 			AppId:    lock.app,
// 			TenantId: lock.tenant,
// 			Resource: lock.resource,
// 			LockId:   l.lockId,
// 		}.Build()
// 	})
// }

// func (l *lock) getRegisteredRLocksForPb() []*managementpb.Lock {
// 	if len(l.rLocks) == 0 {
// 		return nil
// 	}

// 	return lo.Map(l.rLocks, func(lock lock, _ int) *managementpb.Lock {
// 		return managementpb.Lock_builder{
// 			AppId:    lock.app,
// 			TenantId: lock.tenant,
// 			Resource: lock.resource,
// 			LockId:   l.lockId,
// 		}.Build()
// 	})
// }

func (l *lock) Renew() error {
	l.mx.Lock()
	defer l.mx.Unlock()

	ctx, cancel := context.WithCancel(context.Background())
	l.cancelLock = cancel

	lockResponse, err := l.client.Lock(l.ctx, managementpb.LockRequest_builder{
		App:      l.appPrefix,
		Tenant:   l.tenant,
		Ip:       l.ownIp,
		Ttl:      l.ttl,
		Resource: l.resource,
		LockId:   l.lockId,
		// AlreadyRegisteredLocks:  l.getRegisteredLocksForPb(),
		// AlreadyRegisteredRlocks: l.getRegisteredRLocksForPb(),
	}.Build())
	if err != nil {
		return err
	}

	l.lockId = lockResponse.GetLockId()

	go l.keepLockAlive(ctx)

	l.logger.Debug().Write("renewed lock")
	return nil
}

func (l *lock) keepLockAlive(lockCtx context.Context) {
	ticker := time.NewTicker(time.Duration(l.ttl/2) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-lockCtx.Done():
			l.mx.Lock()
			l.cancelLockIfExists()
			l.mx.Unlock()
			return
		case <-l.ctx.Done():
			_, err := l.client.Unlock(context.Background(), managementpb.UnlockRequest_builder{
				App:      l.appPrefix,
				Tenant:   l.tenant,
				Resource: l.resource,
				LockId:   l.lockId,
			}.Build())
			if err != nil {
				l.logger.Warn().WithError(err).Write("unable to drop lock")
			}
			l.stopWg.Done()
			l.mx.Lock()
			l.cancelLockIfExists()
			l.mx.Unlock()
			return
		case <-ticker.C:
			_, err := l.client.ExtendTTL(lockCtx, managementpb.ExtendTTLRequest_builder{
				App:      l.appPrefix,
				Tenant:   l.tenant,
				Resource: l.resource,
				LockId:   l.lockId,
				Ttl:      l.ttl,
			}.Build())
			if err != nil {
				l.logger.Warn().WithError(err).Write("unable to keep lock")

				l.mx.Lock()
				l.onRenew()
				l.cancelLockIfExists()
				l.mx.Unlock()

				if err := util.Retry(l.Renew, time.Millisecond*100, 50); err != nil {
					l.logger.Fatal().WithError(err).Write("unable to renew lock")
				}
				return
			}
		}
	}
}

func (l *lock) cancelLockIfExists() {
	if l.cancelLock != nil {
		l.cancelLock()
		l.cancelLock = nil
	}
}

// func (l *lock) Track(tenant string, resource []string, read bool) {
// 	l.mx.Lock()
// 	defer l.mx.Unlock()

// 	if read {
// 		l.rLocks = append(l.rLocks, lock{
// 			tenant:   tenant,
// 			resource: resource,
// 		})
// 		return
// 	}

// 	l.locks = append(l.locks, lock{
// 		tenant:   tenant,
// 		resource: resource,
// 	})
// }

// func (l *lock) Untrack(tenant string, resource []string, read bool) {
// 	l.mx.Lock()
// 	defer l.mx.Unlock()

// 	if read {
// 		_, index, found := lo.FindIndexOf(l.rLocks, func(rLock lock) bool {
// 			return rLock.tenant == tenant && util.SlicesEqual(rLock.resource, resource)
// 		})

// 		if found {
// 			l.rLocks[index] = l.rLocks[len(l.rLocks)-1]
// 			l.rLocks = l.rLocks[:len(l.rLocks)-1]
// 		}
// 		return
// 	}

// 	l.locks = lo.Filter(l.locks, func(lock lock, _ int) bool {
// 		return lock.tenant != tenant || !util.SlicesEqual(lock.resource, resource)
// 	})
// }
