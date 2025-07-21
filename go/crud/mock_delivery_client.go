package crud

import (
	"context"

	"github.com/fraym/freym-api/go/proto/crud/deliverypb"
	"github.com/stretchr/testify/mock"
)

type MockDeliveryClient struct {
	mock.Mock
}

func (c *MockDeliveryClient) GetEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	filter *Filter,
	wait *Wait,
	returnEmptyDataIfNotFound bool,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*Entry, error) {
	args := c.Called(ctx, typeName, authData, id, filter, returnEmptyDataIfNotFound, wait, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Entry), args.Error(1)
}

func (c *MockDeliveryClient) GetEntryList(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
	wait *ListWait,
) (*EntryList, error) {
	args := c.Called(ctx, typeName, authData, pagination, filter, order, useStrongConsistency, target, wait)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*EntryList), args.Error(1)
}

func (c *MockDeliveryClient) GetViewEntry(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *Filter,
	wait *Wait,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*Entry, error) {
	args := c.Called(ctx, view, authData, filter, wait, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Entry), args.Error(1)
}

func (c *MockDeliveryClient) GetViewEntryList(
	ctx context.Context,
	view string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
	wait *ListWait,
) (*EntryList, error) {
	args := c.Called(ctx, view, authData, pagination, filter, order, useStrongConsistency, target, wait)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*EntryList), args.Error(1)
}

func (c *MockDeliveryClient) CreateEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	data Entry,
	eventMetadata *EventMetadata,
) (*CreateResponse, error) {
	args := c.Called(ctx, typeName, authData, id, data, eventMetadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*CreateResponse), args.Error(1)
}

func (c *MockDeliveryClient) UpsertEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	data Entry,
	eventMetadata *EventMetadata,
) (*UpsertResponse, error) {
	args := c.Called(ctx, typeName, authData, id, data, eventMetadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*UpsertResponse), args.Error(1)
}

func (c *MockDeliveryClient) UpdateEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	data Entry,
	eventMetadata *EventMetadata,
) (*UpdateResponse, error) {
	args := c.Called(ctx, typeName, authData, id, data, eventMetadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*UpdateResponse), args.Error(1)
}

func (c *MockDeliveryClient) UpdateEntryByFilter(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	filter *Filter,
	data Entry,
	eventMetadata *EventMetadata,
) (*UpdateByFilterResponse, error) {
	args := c.Called(ctx, typeName, authData, filter, data, eventMetadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*UpdateByFilterResponse), args.Error(1)
}

func (c *MockDeliveryClient) CloneEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	newId string,
	eventMetadata *EventMetadata,
) (*CloneResponse, error) {
	args := c.Called(ctx, typeName, authData, id, newId, eventMetadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*CloneResponse), args.Error(1)
}

func (c *MockDeliveryClient) DeleteEntryById(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	eventMetadata *EventMetadata,
) (int64, error) {
	args := c.Called(ctx, typeName, authData, id, eventMetadata)
	return args.Get(0).(int64), args.Error(1)
}

func (c *MockDeliveryClient) DeleteEntryByFilter(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	filter *Filter,
	eventMetadata *EventMetadata,
) (int64, error) {
	args := c.Called(ctx, typeName, authData, filter, eventMetadata)
	return args.Get(0).(int64), args.Error(1)
}

func (c *MockDeliveryClient) Close() error {
	return c.Called().Error(0)
}
