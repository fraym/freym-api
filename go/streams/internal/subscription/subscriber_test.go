package subscription_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/freym-api/go/streams/config"
	"github.com/fraym/freym-api/go/streams/domain/dto"
	"github.com/fraym/freym-api/go/streams/internal/subscription"
	"github.com/fraym/freym-api/go/streams/internal/subscription/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubscriber_Ack(t *testing.T) {
	config := config.NewDefaultConfig()
	config.GroupId = "groupId"
	connection := mocks.NewMockConnection()
	endpoint, handler := mocks.NewMockEndpoint(connection)

	expectedTenantId := "test-tenant"
	expectedTopic := "test-topic"
	expectedIncludedTopics := []string{expectedTopic}
	expectedError := ""
	expectedEventId := "eventId"

	wg := sync.WaitGroup{}
	wg.Add(1)

	subscriber := subscription.NewSubscriber(context.Background(), connection, endpoint, config, func(err error) {})
	ctx, cancel := context.WithCancel(context.Background())

	connection.On("Subscribe", ctx, config.GroupId, expectedIncludedTopics).
		Return(nil).
		Run(func(args mock.Arguments) { wg.Done() })
	connection.On("EventHandled", mock.Anything, expectedTenantId, expectedTopic, expectedError).
		Return(nil).
		Run(func(args mock.Arguments) { wg.Done() })

	go func() {
		err := subscriber.Subscribe(
			ctx,
			expectedIncludedTopics,
			func(handleCtx context.Context, event *dto.SubscriptionEvent) error {
				assert.Equal(t, expectedTenantId, event.TenantId)
				assert.Equal(t, expectedTopic, event.Topic)
				assert.Equal(t, expectedEventId, event.Id)

				defer cancel()

				return nil
			},
		)
		assert.NoError(t, err)

		// test receiving event on that stream
		wg.Add(1)
		go func() {
			_ = handler(managementpb.SubscribeResponse_builder{
				Event: managementpb.Event_builder{
					TenantId: expectedTenantId,
					Topic:    expectedTopic,
					Id:       expectedEventId,
					RaisedAt: time.Now().UnixNano(),
				}.Build(),
			}.Build())
		}()
	}()

	wg.Wait()

	_ = handler(managementpb.SubscribeResponse_builder{
		Subscribed: managementpb.Subscribed_builder{
			Error: "",
		}.Build(),
	}.Build())

	<-ctx.Done()
	wg.Wait()
	connection.AssertExpectations(t)
}

func TestSubscriber_NotAck(t *testing.T) {
	config := config.NewDefaultConfig()
	config.GroupId = "groupId"
	connection := mocks.NewMockConnection()
	endpoint, handler := mocks.NewMockEndpoint(connection)

	expectedIncludedTopics := []string{"test-topic"}

	wg := sync.WaitGroup{}
	wg.Add(1)

	subscriber := subscription.NewSubscriber(context.Background(), connection, endpoint, config, func(err error) {})
	ctx, cancel := context.WithCancel(context.Background())

	connection.On("Subscribe", ctx, config.GroupId, expectedIncludedTopics).
		Return(nil).
		Run(func(args mock.Arguments) { wg.Done() })

	go func() {
		err := subscriber.Subscribe(
			ctx,
			expectedIncludedTopics,
			func(handleCtx context.Context, event *dto.SubscriptionEvent) error {
				return errors.New("test")
			},
		)
		assert.Error(t, err)
		cancel()
	}()

	wg.Wait()

	_ = handler(managementpb.SubscribeResponse_builder{
		Subscribed: managementpb.Subscribed_builder{
			Error: "test",
		}.Build(),
	}.Build())

	<-ctx.Done()
	connection.AssertExpectations(t)
}

func TestSubscriber_Timeout(t *testing.T) {
	config := config.NewDefaultConfig()
	config.GroupId = "groupId"
	connection := mocks.NewMockConnection()
	endpoint, _ := mocks.NewMockEndpoint(connection)

	expectedIncludedTopics := []string{"test-topic"}

	wg := sync.WaitGroup{}
	wg.Add(1)

	subscriber := subscription.NewSubscriber(context.Background(), connection, endpoint, config, func(err error) {})
	ctx, cancel := context.WithCancel(context.Background())

	connection.On("Subscribe", ctx, config.GroupId, expectedIncludedTopics).
		Return(nil).
		Run(func(args mock.Arguments) { wg.Done() })

	go func() {
		err := subscriber.Subscribe(
			ctx,
			expectedIncludedTopics,
			func(handleCtx context.Context, event *dto.SubscriptionEvent) error {
				<-handleCtx.Done()
				return handleCtx.Err()
			},
		)
		assert.Error(t, err)
		cancel()
	}()

	wg.Wait()

	<-ctx.Done()
	connection.AssertExpectations(t)
}
