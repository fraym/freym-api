package projections

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockDeliveryClient struct {
	mock.Mock
}

func (c *MockDeliveryClient) GetData(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	filter *Filter,
	returnEmptyDataIfNotFound bool,
	wait *Wait,
	useStrongConsistency bool,
) (*Data, error) {
	args := c.Called(ctx, projection, authData, id, filter, returnEmptyDataIfNotFound, wait, useStrongConsistency)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Data), args.Error(1)
}

func (c *MockDeliveryClient) GetViewData(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *Filter,
) (*Data, error) {
	args := c.Called(ctx, view, authData, filter)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Data), args.Error(1)
}

func (c *MockDeliveryClient) GetDataList(
	ctx context.Context,
	projection string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
) (*DataList, error) {
	args := c.Called(ctx, projection, authData, pagination, filter, order, useStrongConsistency)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*DataList), args.Error(1)
}

func (c *MockDeliveryClient) GetViewDataList(
	ctx context.Context,
	view string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
) (*DataList, error) {
	args := c.Called(ctx, view, authData, pagination, filter, order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*DataList), args.Error(1)
}

func (c *MockDeliveryClient) UpsertData(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	payload Data,
	eventMetadata *EventMetadata,
) (*UpsertResponse, error) {
	args := c.Called(ctx, projection, authData, id, payload, eventMetadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*UpsertResponse), args.Error(1)
}

func (c *MockDeliveryClient) DeleteDataById(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	eventMetadata *EventMetadata,
) (int64, error) {
	args := c.Called(ctx, projection, authData, id, eventMetadata)
	return args.Get(0).(int64), args.Error(1)
}

func (c *MockDeliveryClient) DeleteDataByFilter(
	ctx context.Context,
	projection string,
	authData *AuthData,
	filter *Filter,
	eventMetadata *EventMetadata,
) (int64, error) {
	args := c.Called(ctx, projection, authData, filter, eventMetadata)
	return args.Get(0).(int64), args.Error(1)
}

func (c *MockDeliveryClient) Close() error {
	return c.Called().Error(0)
}
