package mocks

import (
	"sync"

	"github.com/fraym/freym-api/go/streams/internal/subscription"
	"github.com/fraym/golog"
	"github.com/stretchr/testify/mock"
)

func NewMockEndpoint(connection *mockConnection) (*subscription.Endpoint, subscription.HandlerFn) {
	logger := golog.NewLogrusLogger()
	e := subscription.NewEndpoint(logger)

	wg := sync.WaitGroup{}
	wg.Add(1)
	var handler subscription.HandlerFn

	connection.On("Handle", mock.Anything).Return().Run(func(args mock.Arguments) {
		handler = args[0].(subscription.HandlerFn)
		wg.Done()
	})

	go func() {
		_ = e.UseConnection(connection)
	}()

	wg.Wait()

	return e, handler
}
