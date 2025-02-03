package client_test

import (
	"context"
	"testing"

	"github.com/fraym/freym-api/go/streams/domain/dto"
	"github.com/fraym/freym-api/go/streams/internal/client"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockEventHandler struct {
	mock.Mock
}

func (h *mockEventHandler) Handle(ctx context.Context, event *dto.SubscriptionEvent) error {
	return h.Called(ctx, event).Error(0)
}

func TestHandler_NoHandler(t *testing.T) {
	handler := client.NewHandler(false)

	expectedType := "test-type"

	err := handler.Handle(context.Background(), &dto.SubscriptionEvent{
		Type: expectedType,
	})
	assert.Error(t, err)
}

func TestHandler_NoHandlerIgnoreUnhandled(t *testing.T) {
	handler := client.NewHandler(true)

	expectedType := "test-type"

	err := handler.Handle(context.Background(), &dto.SubscriptionEvent{
		Type: expectedType,
	})
	assert.NoError(t, err)
}

func TestHandler_Handler(t *testing.T) {
	handler := client.NewHandler(false)
	eventHandler := &mockEventHandler{}

	expectedType := "test-type"
	expectedEvent := &dto.SubscriptionEvent{
		Type: expectedType,
	}

	eventHandler.On("Handle", mock.Anything, expectedEvent).Return(nil)

	handler.Use(expectedType, eventHandler.Handle)
	err := handler.Handle(context.Background(), expectedEvent)
	assert.NoError(t, err)

	eventHandler.AssertExpectations(t)
}

func TestHandler_Handler_Err(t *testing.T) {
	handler := client.NewHandler(false)
	eventHandler := &mockEventHandler{}

	expectedType := "test-type"
	expectedEvent := &dto.SubscriptionEvent{
		Type: expectedType,
	}

	eventHandler.On("Handle", mock.Anything, expectedEvent).Return(errors.New("test"))

	handler.Use(expectedType, eventHandler.Handle)
	err := handler.Handle(context.Background(), expectedEvent)
	assert.Error(t, err)

	eventHandler.AssertExpectations(t)
}

func TestHandler_Handler_AllTypes(t *testing.T) {
	handler := client.NewHandler(false)
	allEventsHandler := &mockEventHandler{}
	eventHandler := &mockEventHandler{}

	expectedType := "test-type"
	expectedEvent := &dto.SubscriptionEvent{
		Type: expectedType,
	}

	allEventsHandler.On("Handle", mock.Anything, expectedEvent).Return(nil)
	eventHandler.On("Handle", mock.Anything, expectedEvent).Return(nil)

	handler.UseForAllEventTypes(allEventsHandler.Handle)
	handler.Use(expectedType, eventHandler.Handle)
	err := handler.Handle(context.Background(), expectedEvent)
	assert.NoError(t, err)

	allEventsHandler.AssertExpectations(t)
	eventHandler.AssertExpectations(t)
}
