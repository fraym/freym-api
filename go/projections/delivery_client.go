package projections

import (
	"context"
	"time"

	"github.com/Becklyn/go-wire-core/json"
	"github.com/fraym/freym-api/go/projections/config"
	"github.com/fraym/freym-api/go/projections/delivery"
	"github.com/fraym/freym-api/go/proto/projections/deliverypb"
)

type DataList struct {
	Limit int64
	Page  int64
	Total int64
	Data  []Data
}

type Data map[string]any

type UpsertResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
	Data                  Data
	ID                    string
}

type Wait struct {
	ConditionFilter  *Filter
	ConditionTimeout time.Duration
}

func (w *Wait) toDeliveryWait() *deliverypb.DataWait {
	if w == nil || w.ConditionFilter == nil {
		return nil
	}

	return deliverypb.DataWait_builder{
		ConditionFilter: w.ConditionFilter.toProtobufFilter(),
		Timeout:         w.ConditionTimeout.Milliseconds(),
	}.Build()
}

type DeliveryClient interface {
	GetData(
		ctx context.Context,
		projection string,
		authData *AuthData,
		id string,
		filter *Filter,
		returnEmptyDataIfNotFound bool,
		wait *Wait,
		useStrongConsistency bool,
		target deliverypb.DeploymentTarget,
	) (*Data, error)
	GetViewData(
		ctx context.Context,
		view string,
		authData *AuthData,
		filter *Filter,
		useStrongConsistency bool,
		target deliverypb.DeploymentTarget,
	) (*Data, error)
	GetDataList(
		ctx context.Context,
		projection string,
		authData *AuthData,
		pagination *Pagination,
		filter *Filter,
		order []Order,
		useStrongConsistency bool,
		target deliverypb.DeploymentTarget,
	) (*DataList, error)
	GetViewDataList(
		ctx context.Context,
		view string,
		authData *AuthData,
		pagination *Pagination,
		filter *Filter,
		order []Order,
		useStrongConsistency bool,
		target deliverypb.DeploymentTarget,
	) (*DataList, error)
	UpsertData(
		ctx context.Context,
		projection string,
		authData *AuthData,
		id string,
		payload Data,
		eventMetadata *EventMetadata,
	) (*UpsertResponse, error)
	DeleteDataById(
		ctx context.Context,
		projection string,
		authData *AuthData,
		id string,
		eventMetadata *EventMetadata,
	) (int64, error)
	DeleteDataByFilter(
		ctx context.Context,
		projection string,
		authData *AuthData,
		filter *Filter,
		eventMetadata *EventMetadata,
	) (int64, error)
	Close() error
}

type projectionsDeliveryClient struct {
	clientCloseFn func() error
	client        deliverypb.ServiceClient
}

func NewDeliveryClient(
	conf *config.Config,
) (DeliveryClient, error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	client, clientCloseFn, err := delivery.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &projectionsDeliveryClient{
		clientCloseFn: clientCloseFn,
		client:        client,
	}, nil
}

func (c *projectionsDeliveryClient) GetData(
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
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetData(ctx, deliverypb.GetDataRequest_builder{
		Projection:                projection,
		Auth:                      pbAuthData,
		DataId:                    id,
		Filter:                    filter.toProtobufFilter(),
		ReturnEmptyDataIfNotFound: returnEmptyDataIfNotFound,
		Wait:                      wait.toDeliveryWait(),
		UseStrongConsistency:      useStrongConsistency,
		Target:                    target,
	}.Build())
	if err != nil {
		return nil, err
	}

	result := response.GetResult()

	if result == nil {
		return nil, nil
	}
	return getDataMap(result.GetData())
}

func (c *projectionsDeliveryClient) GetViewData(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *Filter,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*Data, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetViewData(ctx, deliverypb.GetViewDataRequest_builder{
		View:                 view,
		Auth:                 pbAuthData,
		Filter:               filter.toProtobufFilter(),
		UseStrongConsistency: useStrongConsistency,
		Target:               target,
	}.Build())
	if err != nil {
		return nil, err
	}

	result := response.GetResult()

	if result == nil {
		return nil, nil
	}
	return getDataMap(result.GetData())
}

func (c *projectionsDeliveryClient) GetDataList(
	ctx context.Context,
	projection string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*DataList, error) {
	var (
		limit int64
		page  int64
	)

	if pagination != nil {
		limit = pagination.Limit
		page = pagination.Page
	}

	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetDataList(ctx, deliverypb.GetDataListRequest_builder{
		Projection:           projection,
		Auth:                 pbAuthData,
		Limit:                limit,
		Page:                 page,
		Filter:               filter.toProtobufFilter(),
		Order:                toProtobufOrder(order),
		UseStrongConsistency: useStrongConsistency,
		Target:               target,
	}.Build())
	if err != nil {
		return nil, err
	}

	var result []Data

	for _, element := range response.GetResult() {
		data, err := getDataMap(element.GetData())
		if err != nil {
			return nil, err
		}

		result = append(result, *data)
	}

	return &DataList{
		Limit: response.GetLimit(),
		Page:  response.GetPage(),
		Total: response.GetTotal(),
		Data:  result,
	}, nil
}

func (c *projectionsDeliveryClient) GetViewDataList(
	ctx context.Context,
	view string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
) (*DataList, error) {
	var (
		limit int64
		page  int64
	)

	if pagination != nil {
		limit = pagination.Limit
		page = pagination.Page
	}

	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetViewDataList(ctx, deliverypb.GetViewDataListRequest_builder{
		View:                 view,
		Auth:                 pbAuthData,
		Limit:                limit,
		Page:                 page,
		Filter:               filter.toProtobufFilter(),
		Order:                toProtobufOrder(order),
		UseStrongConsistency: useStrongConsistency,
		Target:               target,
	}.Build())
	if err != nil {
		return nil, err
	}

	var result []Data

	for _, element := range response.GetResult() {
		data, err := getDataMap(element.GetData())
		if err != nil {
			return nil, err
		}

		result = append(result, *data)
	}

	return &DataList{
		Limit: response.GetLimit(),
		Page:  response.GetPage(),
		Total: response.GetTotal(),
		Data:  result,
	}, nil
}

func (c *projectionsDeliveryClient) UpsertData(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	payload Data,
	eventMetadata *EventMetadata,
) (*UpsertResponse, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	pbPayload := map[string]string{}

	for key, value := range payload {
		byteValue, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}

		pbPayload[key] = string(byteValue)
	}

	response, err := c.client.Upsert(ctx, deliverypb.UpsertRequest_builder{
		Projection:    projection,
		Auth:          pbAuthData,
		DataId:        id,
		Payload:       pbPayload,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return nil, err
	}

	data, err := getDataMap(response.GetNewData())
	if err != nil {
		return nil, err
	}

	return &UpsertResponse{
		ValidationErrors:      response.GetValidationErrors(),
		FieldValidationErrors: response.GetFieldValidationErrors(),
		Data:                  *data,
		ID:                    response.GetId(),
	}, nil
}

func (c *projectionsDeliveryClient) DeleteDataById(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	eventMetadata *EventMetadata,
) (int64, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return 0, err
	}

	response, err := c.client.Delete(ctx, deliverypb.DeleteRequest_builder{
		Projection:    projection,
		Auth:          pbAuthData,
		DataId:        id,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return 0, err
	}

	return response.GetNumberOfDeletedEntries(), nil
}

func (c *projectionsDeliveryClient) DeleteDataByFilter(
	ctx context.Context,
	projection string,
	authData *AuthData,
	filter *Filter,
	eventMetadata *EventMetadata,
) (int64, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return 0, err
	}

	response, err := c.client.Delete(ctx, deliverypb.DeleteRequest_builder{
		Projection:    projection,
		Auth:          pbAuthData,
		Filter:        filter.toProtobufFilter(),
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return 0, err
	}

	return response.GetNumberOfDeletedEntries(), nil
}

func (c *projectionsDeliveryClient) Close() error {
	return c.clientCloseFn()
}

func getDataMap(data map[string]string) (*Data, error) {
	dataMap := Data{}

	for key, value := range data {
		var originalData any

		if err := json.Unmarshal([]byte(value), &originalData); err != nil {
			return nil, err
		}

		dataMap[key] = originalData
	}

	return &dataMap, nil
}
