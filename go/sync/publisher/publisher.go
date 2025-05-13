package publisher

import (
	"context"
	"time"

	"github.com/fraym/freym-api/go/sync/internal/peer"
	"github.com/fraym/golog"
)

type SendResponse struct {
	Ack    bool
	Reason *string
}

type SendFn[T peer.ServiceClient, RequestData any] func(ctx context.Context, data RequestData, client T) (*SendResponse, error)

type Publisher[T peer.ServiceClient, RequestData any] struct {
	logger   golog.Logger
	peerPool peer.PeerPool[T]
	send     SendFn[T, RequestData]
}

func NewPublisher[T peer.ServiceClient, RequestData any](
	logger golog.Logger,
	peerPool peer.PeerPool[T],
	send SendFn[T, RequestData],
) *Publisher[T, RequestData] {
	return &Publisher[T, RequestData]{
		logger:   logger,
		peerPool: peerPool,
		send:     send,
	}
}

func (p *Publisher[T, RequestData]) waitForRetryPause(numberOfRetry int) {
	switch numberOfRetry {
	case 1:
		time.Sleep(100 * time.Millisecond)
	case 2:
		time.Sleep(500 * time.Millisecond)
	}
}

// trySendToClient returns true if the client acknowledged the request
func (p *Publisher[T, RequestData]) TrySendToClient(
	data RequestData,
	address string,
	client T,
	numberOfRetry int,
) *Response {
	p.waitForRetryPause(numberOfRetry)

	if numberOfRetry == 3 {
		p.logger.Warn().WithFields(map[string]any{
			"address": address,
			"attempt": numberOfRetry,
		}).Write("stopped retrying to send backchannel message")
		return NewResponse(false, false)
	}

	if numberOfRetry > 0 && !p.peerPool.IsPeer(address) {
		p.logger.Warn().WithFields(map[string]any{
			"address": address,
			"attempt": numberOfRetry,
		}).Write("receiving client left the backchannel")
		return NewResponse(false, true)
	}

	response, err := p.send(context.Background(), data, client)
	if err != nil {
		p.logger.Warn().WithError(err).WithFields(map[string]any{
			"address": address,
			"attempt": numberOfRetry,
		}).Write("failed to send backchannel to client")

		return p.TrySendToClient(data, address, client, numberOfRetry+1)
	}

	if response.Ack {
		return NewResponse(true, false)
	}

	if response.Reason != nil {
		p.logger.Debug().WithFields(map[string]any{
			"address": address,
			"attempt": numberOfRetry,
			"reason":  *response.Reason,
		}).Write("client did not acknowledge the backchannel message")
	} else {
		p.logger.Debug().WithFields(map[string]any{
			"address": address,
			"attempt": numberOfRetry,
		}).Write("client did not acknowledge the backchannel message")
	}

	return NewResponse(false, false)
}
