package client

import (
	"context"
	"time"

	"github.com/fraym/freym-api/go/proto/sync/managementpb"
	"github.com/fraym/freym-api/go/sync/internal/util"
	"github.com/fraym/golog"
)

type Service struct {
	connection *connection
	lock       *lock
	peers      *peers
	client     managementpb.ServiceClient
	logger     golog.Logger
	retryPause time.Duration
}

func NewService(
	ctx context.Context,
	client managementpb.ServiceClient,
	logger golog.Logger,
	appPrefix string,
	ownIp string,
) (*Service, error) {
	retryPause := time.Millisecond * 10000
	connection := NewConnection(ctx)
	lock := NewLock(ctx, appPrefix, ownIp, client, logger)

	connection.OnConnect(func() {
		if err := util.Retry(lock.Renew, retryPause, 10); err != nil {
			logger.Fatal().WithError(err).Write("unable to renew lease")
		}
	})

	peers, err := NewPeers(ctx, client, logger, appPrefix, connection.Connect, connection.Disconnect)
	if err != nil {
		return nil, err
	}

	lock.OnRenew(func() {
		connection.Disconnect(len(peers.PeerIps()))
		peers.RenewConnection()
	})

	peers.OnPeerUpdate(func(peerIps []string) {
		connection.PeerUpdate(len(peerIps))
	})

	return &Service{
		connection: connection,
		lock:       lock,
		peers:      peers,
		client:     client,
		logger:     logger,
		retryPause: retryPause,
	}, nil
}

func (s *Service) Lock(tenant string, resource []string) error {
	if err := s.connection.WaitForConnect(); err != nil {
		return err
	}

	if err := util.Retry(func() error {
		_, err := s.client.Lock(context.Background(), managementpb.LockRequest_builder{
			LockId:   s.lock.lockId,
			App:      s.lock.appPrefix,
			Tenant:   tenant,
			Ip:       s.lock.ownIp,
			Ttl:      s.lock.ttl,
			Resource: resource,
		}.Build())
		return err
	}, s.retryPause, 50); err != nil {
		return err
	}

	return nil
}

func (s *Service) Unlock(tenant string, resource []string) {
	go func() {
		if err := s.connection.WaitForConnect(); err != nil {
			s.logger.Fatal().WithFields(map[string]any{
				"tenant":   tenant,
				"resource": resource,
			}).WithError(err).Write("unable to unlock")
			return
		}

		if err := util.Retry(func() error {
			_, err := s.client.Unlock(context.Background(), managementpb.UnlockRequest_builder{
				LockId:   s.lock.lockId,
				App:      s.lock.appPrefix,
				Tenant:   tenant,
				Resource: resource,
			}.Build())
			return err
		}, s.retryPause, 50); err != nil {
			s.logger.Fatal().WithFields(map[string]any{
				"tenant":   tenant,
				"resource": resource,
			}).WithError(err).Write("unable to unlock")
		}
	}()
}

func (s *Service) RLock(tenant string, resource []string) error {
	if err := s.connection.WaitForConnect(); err != nil {
		return err
	}

	if err := util.Retry(func() error {
		_, err := s.client.RLock(context.Background(), managementpb.RLockRequest_builder{
			LockId:   s.lock.lockId,
			App:      s.lock.appPrefix,
			Tenant:   tenant,
			Ip:       s.lock.ownIp,
			Ttl:      s.lock.ttl,
			Resource: resource,
		}.Build())
		return err
	}, s.retryPause, 50); err != nil {
		return err
	}

	return nil
}

func (s *Service) RUnlock(tenant string, resource []string) {
	go func() {
		if err := s.connection.WaitForConnect(); err != nil {
			s.logger.Fatal().WithFields(map[string]any{
				"tenant":   tenant,
				"resource": resource,
			}).WithError(err).Write("unable to runlock")
			return
		}

		if err := util.Retry(func() error {
			_, err := s.client.RUnlock(context.Background(), managementpb.RUnlockRequest_builder{
				LockId:   s.lock.lockId,
				App:      s.lock.appPrefix,
				Tenant:   tenant,
				Resource: resource,
			}.Build())
			return err
		}, s.retryPause, 50); err != nil {
			s.logger.Fatal().WithFields(map[string]any{
				"tenant":   tenant,
				"resource": resource,
			}).WithError(err).Write("unable to runlock")
		}
	}()
}

func (s *Service) WaitForStop() {
	s.lock.WaitForStop()
}

func (s *Service) IsPeer(address string) bool {
	return s.peers.IsPeer(address)
}

func (s *Service) PeerIps() []string {
	return s.peers.PeerIps()
}

func (s *Service) OnPeerLeave(handler func(peerIp string)) {
	s.peers.OnPeerLeave(handler)
}
