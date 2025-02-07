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

type peers struct {
	ctx          context.Context
	client       managementpb.ServiceClient
	logger       golog.Logger
	appPrefix    string
	onConnect    func()
	onDisconnect func(numberOfPeerNodes int)

	mx           *sync.RWMutex
	peerIps      []string
	onPeerLeave  func(peerIp string)
	onPeerUpdate func(peerIps []string)

	connectionCtx    context.Context
	cancelConnection context.CancelFunc
}

func NewPeers(
	ctx context.Context,
	client managementpb.ServiceClient,
	logger golog.Logger,
	appPrefix string,
	onConnect func(),
	onDisconnect func(numberOfPeerNodes int),
) (*peers, error) {
	p := &peers{
		ctx:          ctx,
		client:       client,
		logger:       logger,
		appPrefix:    appPrefix,
		onConnect:    onConnect,
		onDisconnect: onDisconnect,

		mx:      &sync.RWMutex{},
		peerIps: nil,
	}

	return p, p.init()
}

func (p *peers) init() error {
	p.connectionCtx, p.cancelConnection = context.WithCancel(p.ctx)

	service, err := p.client.GetPeerNodes(p.connectionCtx, managementpb.GetPeerNodesRequest_builder{
		App: p.appPrefix,
	}.Build())
	if err != nil {
		return err
	}

	list, err := service.Recv()
	if err != nil {
		return err
	}

	p.mx.Lock()
	p.peerIps = list.GetPeerNodeIp()
	if p.onPeerUpdate != nil {
		p.onPeerUpdate(p.peerIps)
	}
	p.mx.Unlock()

	p.onConnect()

	go p.watchPeerNodes(service)

	return nil
}

func (p *peers) watchPeerNodes(service managementpb.Service_GetPeerNodesClient) {
	for {
		list, err := service.Recv()
		if err != nil {
			select {
			case <-p.ctx.Done():
				return
			default:
			}

			p.mx.RLock()
			p.onDisconnect(len(p.peerIps))
			p.mx.RUnlock()

			if err := util.Retry(func() error {
				return p.init()
			}, time.Millisecond*100, 50); err != nil {
				p.logger.Fatal().WithError(err).Write("unable to reconnect to sync service")
			}
			return
		}

		p.mx.Lock()

		if p.onPeerLeave != nil {
			diff, _ := lo.Difference(p.peerIps, list.GetPeerNodeIp())

			for _, peerIp := range lo.Uniq(diff) {
				p.onPeerLeave(peerIp)
			}
		}

		p.peerIps = list.GetPeerNodeIp()

		if p.onPeerUpdate != nil {
			p.onPeerUpdate(p.peerIps)
		}
		p.mx.Unlock()
	}
}

func (p *peers) PeerIps() []string {
	p.mx.RLock()
	defer p.mx.RUnlock()

	return p.peerIps
}

func (p *peers) IsPeer(address string) bool {
	p.mx.RLock()
	defer p.mx.RUnlock()

	return lo.Contains(p.peerIps, address)
}

func (p *peers) OnPeerLeave(handler func(peerIp string)) {
	p.mx.Lock()
	defer p.mx.Unlock()

	p.onPeerLeave = handler
}

func (p *peers) OnPeerUpdate(handler func(peerIps []string)) {
	p.mx.Lock()
	defer p.mx.Unlock()

	p.onPeerUpdate = handler
}

func (p *peers) RenewConnection() {
	p.cancelConnection()
}
