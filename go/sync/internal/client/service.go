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
	lease      *lease
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
	retryPause := time.Millisecond * 1000
	connection := NewConnection(ctx)

	lease := NewLease(ctx, appPrefix, ownIp, client, logger)

	connection.OnConnect(func() {
		if err := util.Retry(lease.Renew, retryPause, 10); err != nil {
			logger.Fatal().WithError(err).Write("unable to renew lease")
		}
	})

	peers, err := NewPeers(ctx, client, logger, appPrefix, connection.Connect, connection.Disconnect)
	if err != nil {
		return nil, err
	}

	lease.OnRenew(func() {
		connection.Disconnect(len(peers.PeerIps()))
		peers.RenewConnection()
	})

	peers.OnPeerUpdate(func(peerIps []string) {
		connection.PeerUpdate(len(peerIps))
	})

	return &Service{
		connection: connection,
		lease:      lease,
		peers:      peers,
		client:     client,
		logger:     logger,
		retryPause: retryPause,
	}, nil
}

func (s *Service) Lock(ttl int32, tenant string, resource []string) error {
	if err := s.connection.WaitForConnect(); err != nil {
		return err
	}

	if err := util.Retry(func() error {
		_, err := s.client.Lock(context.Background(), managementpb.LockRequest_builder{
			LeaseId:  s.lease.LeaseId(),
			TenantId: tenant,
			Resource: resource,
			Ttl:      ttl,
		}.Build())
		return err
	}, s.retryPause, 50); err != nil {
		return err
	}

	s.lease.Track(tenant, resource, false)

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
				LeaseId:  s.lease.LeaseId(),
				TenantId: tenant,
				Resource: resource,
			}.Build())
			return err
		}, s.retryPause, 50); err != nil {
			s.logger.Fatal().WithFields(map[string]any{
				"tenant":   tenant,
				"resource": resource,
			}).WithError(err).Write("unable to unlock")
		} else {
			s.lease.Untrack(tenant, resource, false)
		}
	}()
}

func (s *Service) RLock(ttl int32, tenant string, resource []string) error {
	if err := s.connection.WaitForConnect(); err != nil {
		return err
	}

	if err := util.Retry(func() error {
		_, err := s.client.RLock(context.Background(), managementpb.RLockRequest_builder{
			LeaseId:  s.lease.LeaseId(),
			TenantId: tenant,
			Resource: resource,
			Ttl:      ttl,
		}.Build())
		return err
	}, s.retryPause, 50); err != nil {
		return err
	}

	s.lease.Track(tenant, resource, true)

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
				LeaseId:  s.lease.LeaseId(),
				TenantId: tenant,
				Resource: resource,
			}.Build())
			return err
		}, s.retryPause, 50); err != nil {
			s.logger.Fatal().WithFields(map[string]any{
				"tenant":   tenant,
				"resource": resource,
			}).WithError(err).Write("unable to runlock")
		} else {
			s.lease.Untrack(tenant, resource, true)
		}
	}()
}

func (s *Service) WaitForStop() {
	s.lease.WaitForStop()
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
