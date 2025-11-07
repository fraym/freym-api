package subscription

import (
	"context"
	"io"
	"time"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/freym-api/go/streams/config"
	"github.com/fraym/golog"
	"github.com/google/uuid"
)

type HandlerFn func(response *managementpb.SubscribeResponse) error

type Connection interface {
	Handle(handlerFn HandlerFn) error
	Close()
	Subscribe(ctx context.Context, groupId string, topics []string, parallelTopicProcessing bool) error
	EventHandled(
		ctx context.Context,
		tenantId string,
		topic string,
		errorMessage string,
		retry bool,
		stream string,
	) error
}

type clientConnection struct {
	deploymentId   int64
	isClosed       bool
	receiveCtx     context.Context
	receiveCancel  context.CancelFunc
	sendCtx        context.Context
	sendCancel     context.CancelFunc
	sendTimeout    time.Duration
	grpcConnection managementpb.Service_SubscribeClient
	requests       chan *managementpb.SubscribeRequest

	logger golog.Logger
	errs   chan error
}

func NewConnection(
	ctx context.Context,
	deploymentId int64,
	client managementpb.ServiceClient,
	logger golog.Logger,
	config *config.Config,
) (Connection, error) {
	grpcConnection, err := client.Subscribe(ctx)
	if err != nil {
		return nil, err
	}

	receiveCtx, receiveCancel := context.WithCancel(grpcConnection.Context())
	sendCtx, sendCancel := context.WithCancel(grpcConnection.Context())

	return &clientConnection{
		deploymentId:   deploymentId,
		isClosed:       false,
		receiveCtx:     receiveCtx,
		receiveCancel:  receiveCancel,
		sendCtx:        sendCtx,
		sendCancel:     sendCancel,
		sendTimeout:    config.SendTimeout,
		grpcConnection: grpcConnection,
		requests:       make(chan *managementpb.SubscribeRequest),
		logger:         logger,
		errs:           make(chan error),
	}, nil
}

func (c *clientConnection) Handle(handlerFn HandlerFn) error {
	go c.listen(handlerFn)
	go c.send()

	return c.wait()
}

func (c *clientConnection) listen(handlerFn HandlerFn) {
	defer func() {
		c.isClosed = true
		c.sendCancel()
		c.receiveCancel()
		c.logger.Debug().Write("gRPC connection stopped listening")
	}()

	errChan := make(chan error)

loop:
	for {
		response, err := c.grpcConnection.Recv()
		if err != nil {
			if !c.isClosed && err != io.EOF {
				errChan <- err
				return
			}
			close(errChan)
			return
		}

		go func() {
			if err := handlerFn(response); err != nil {
				errChan <- err
				return
			}
		}()

		select {
		case err := <-errChan:
			c.errs <- err
			break loop
		default:
		}
	}
}

func (c *clientConnection) send() {
	defer func() {
		c.isClosed = true
		c.sendCancel()
		c.receiveCancel()
		c.logger.Debug().Write("gRPC connection stopped sending")
	}()

	for {
		select {
		case response := <-c.requests:
			err := c.grpcConnection.Send(response)
			if err != nil {
				if !c.isClosed && err != io.EOF {
					c.errs <- err
					return
				}
				close(c.errs)
				return
			}
		case <-c.receiveCtx.Done():
			c.errs <- c.receiveCtx.Err()
			return
		}
	}
}

func (c *clientConnection) wait() error {
	select {
	case <-c.receiveCtx.Done():
	case <-c.sendCtx.Done():
	}

	if err := <-c.errs; err != nil {
		c.logger.Debug().Writef("gRPC connection to streams service ended with error: %s", err)
		return err
	}
	c.logger.Debug().Write("gRPC connection to streams service ended")
	return nil
}

func (c *clientConnection) Close() {
	c.receiveCancel()
	c.sendCancel()
}

func (c *clientConnection) Subscribe(
	ctx context.Context,
	groupId string,
	topics []string,
	parallelTopicProcessing bool,
) error {
	return c.sendRequest(ctx, managementpb.SubscribeRequest_builder{
		Subscribe: managementpb.Subscribe_builder{
			Metadata: managementpb.SubscribeMetadata_builder{
				Group:                   groupId,
				SubscriberId:            uuid.NewString(),
				DeploymentId:            c.deploymentId,
				ParallelTopicProcessing: parallelTopicProcessing,
			}.Build(),
			Topics: topics,
		}.Build(),
	}.Build())
}

func (c *clientConnection) EventHandled(
	ctx context.Context,
	tenantId string,
	topic string,
	errorMessage string,
	retry bool,
	stream string,
) error {
	return c.sendRequest(ctx, managementpb.SubscribeRequest_builder{
		Handled: managementpb.Handled_builder{
			TenantId: tenantId,
			Topic:    topic,
			Error:    errorMessage,
			Retry:    retry,
			Stream:   stream,
		}.Build(),
	}.Build())
}

func (c *clientConnection) sendRequest(ctx context.Context, request *managementpb.SubscribeRequest) error {
	ctx, cancel := context.WithTimeout(ctx, c.sendTimeout)
	defer cancel()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case c.requests <- request:
		return nil
	}
}
