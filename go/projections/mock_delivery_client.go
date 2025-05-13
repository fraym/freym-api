package projections

import (
	"context"

	"github.com/fraym/freym-api/go/proto/projections/deliverypb"
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
	target deliverypb.DeploymentTarget,
) (*Data, error) {
	args := c.Called(
		ctx,
		projection,
		authData,
		id,
		filter,
		returnEmptyDataIfNotFound,
		wait,
		useStrongConsistency,
		target,
	)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Data), args.Error(1)
}

func (c *MockDeliveryClient) GetJsonData(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	filter *JsonFilter,
	returnEmptyDataIfNotFound bool,
	wait *JsonWait,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*JsonData, error) {
	args := c.Called(
		ctx,
		projection,
		authData,
		id,
		filter,
		returnEmptyDataIfNotFound,
		wait,
		useStrongConsistency,
		target,
	)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*JsonData), args.Error(1)
}

func (c *MockDeliveryClient) GetViewData(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *Filter,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*Data, error) {
	args := c.Called(ctx, view, authData, filter, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Data), args.Error(1)
}

func (c *MockDeliveryClient) GetViewJsonData(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *JsonFilter,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*JsonData, error) {
	args := c.Called(ctx, view, authData, filter, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*JsonData), args.Error(1)
}

func (c *MockDeliveryClient) GetDataList(
	ctx context.Context,
	projection string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*DataList, error) {
	args := c.Called(ctx, projection, authData, pagination, filter, order, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*DataList), args.Error(1)
}

func (c *MockDeliveryClient) GetJsonDataList(
	ctx context.Context,
	projection string,
	authData *AuthData,
	pagination *Pagination,
	filter *JsonFilter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*JsonDataList, error) {
	args := c.Called(ctx, projection, authData, pagination, filter, order, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*JsonDataList), args.Error(1)
}

func (c *MockDeliveryClient) GetViewDataList(
	ctx context.Context,
	view string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*DataList, error) {
	args := c.Called(ctx, view, authData, pagination, filter, order, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*DataList), args.Error(1)
}

func (c *MockDeliveryClient) GetViewJsonDataList(
	ctx context.Context,
	view string,
	authData *AuthData,
	pagination *Pagination,
	filter *JsonFilter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*JsonDataList, error) {
	args := c.Called(ctx, view, authData, pagination, filter, order, useStrongConsistency, target)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*JsonDataList), args.Error(1)
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

func (c *MockDeliveryClient) UpsertJsonData(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	payload JsonData,
	eventMetadata *EventMetadata,
) (*JsonUpsertResponse, error) {
	args := c.Called(ctx, projection, authData, id, payload, eventMetadata)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*JsonUpsertResponse), args.Error(1)
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

func (c *MockDeliveryClient) DeleteDataByJsonFilter(
	ctx context.Context,
	projection string,
	authData *AuthData,
	filter *JsonFilter,
	eventMetadata *EventMetadata,
) (int64, error) {
	args := c.Called(ctx, projection, authData, filter, eventMetadata)
	return args.Get(0).(int64), args.Error(1)
}

func (c *MockDeliveryClient) Close() error {
	return c.Called().Error(0)
}
