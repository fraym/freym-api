package sync

import (
	"context"
	"sync"

	"github.com/fraym/freym-api/go/sync/internal/peer"
	"github.com/fraym/freym-api/go/sync/publisher"
	"github.com/fraym/golog"
)

type ParallelPublisher[T peer.ServiceClient, RequestData any] struct {
	publisher.Publisher[T, RequestData]
	peerPool peer.PeerPool[T]
}

func NewParallelPublisher[T peer.ServiceClient, RequestData any](
	logger golog.Logger,
	peerPool peer.PeerPool[T],
	sendFn publisher.SendFn[T, RequestData],
) ParallelPublisher[T, RequestData] {
	return ParallelPublisher[T, RequestData]{
		Publisher: *publisher.NewPublisher(logger, peerPool, sendFn),
		peerPool:  peerPool,
	}
}

func (p *ParallelPublisher[T, RequestData]) Send(
	ctx context.Context,
	data RequestData,
	filter publisher.AddressFilter,
	callback publisher.SendCallback,
) error {
	wg := sync.WaitGroup{}
	err := p.peerPool.Iterate(ctx, func(address string, client T) bool {
		if filter != nil && !filter(address) {
			return true
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			ack := p.TrySendToClient(data, address, client, 0)

			if callback != nil {
				callback(address, ack)
			}
		}()

		return true
	})

	wg.Wait()
	return err
}
