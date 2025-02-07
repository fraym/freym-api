package client

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrWaitForClosedConnection = errors.New("cannot wait for closed connection")
	ErrWaitTooLong             = errors.New("cannot reconnect to sync service")
)

type connection struct {
	ctx                           context.Context
	mx                            sync.RWMutex
	isConnected                   bool
	numberOfPeerNodesOnDisconnect int
	waitForPeerReconnectCtx       context.Context
	allPeersReconnected           context.CancelFunc

	onConnect []func()
}

func NewConnection(ctx context.Context) *connection {
	return &connection{
		ctx:                           ctx,
		mx:                            sync.RWMutex{},
		isConnected:                   false,
		onConnect:                     nil,
		numberOfPeerNodesOnDisconnect: 0,
		waitForPeerReconnectCtx:       context.Background(),
		allPeersReconnected:           nil,
	}
}

func (c *connection) IsConnected() bool {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.isConnected
}

func (c *connection) WaitForConnect() error {
	if c.IsConnected() {
		return nil
	}

	c.mx.Lock()
	if c.isConnected {
		c.mx.Unlock()
		return nil
	}

	connectChan := make(chan bool)
	c.onConnect = append(c.onConnect, func() {
		connectChan <- true
	})
	c.mx.Unlock()

	fiveSecCtx, stop := context.WithTimeout(context.Background(), time.Second*5)
	defer stop()

	for {
		select {
		case <-c.ctx.Done():
			return ErrWaitForClosedConnection
		case <-fiveSecCtx.Done():
			return ErrWaitTooLong
		case <-connectChan:
			<-c.waitForPeerReconnectCtx.Done()
			return nil
		}
	}
}

func (c *connection) Disconnect(numberOfPeerNodes int) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.isConnected = false
	c.numberOfPeerNodesOnDisconnect = numberOfPeerNodes
	c.waitForPeerReconnectCtx, c.allPeersReconnected = context.WithCancel(context.Background())
}

func (c *connection) Connect() {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.isConnected = true

	for _, onConnect := range c.onConnect {
		onConnect()
	}
	c.onConnect = nil
}

func (c *connection) OnConnect(fn func()) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.onConnect = append(c.onConnect, fn)
}

func (c *connection) PeerUpdate(numberOfPeers int) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	if c.allPeersReconnected == nil || c.numberOfPeerNodesOnDisconnect > numberOfPeers {
		return
	}

	c.allPeersReconnected()
}
