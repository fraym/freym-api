package projections

import (
	"context"

	"github.com/Becklyn/go-wire-core/json"
	"github.com/fraym/freym-api/go/projections/config"
	"github.com/fraym/freym-api/go/projections/delivery"
	"github.com/fraym/freym-api/go/proto/projections/deliverypb"
)

type (
	Data     map[string]any
	JsonData map[string]string
	DataList struct {
		Limit int64
		Page  int64
		Total int64
		Data  []Data
	}
	JsonDataList struct {
		Limit int64
		Page  int64
		Total int64
		Data  []JsonData
	}
	UpsertResponse struct {
		ValidationErrors      []string
		FieldValidationErrors map[string]string
		Data                  Data
		ID                    string
	}
	JsonUpsertResponse struct {
		ValidationErrors      []string
		FieldValidationErrors map[string]string
		Data                  JsonData
		ID                    string
	}
)

type DeliveryClient interface {
	GetData(
		ctx context.Context,
		projection string,
		authData *AuthData,
		id string,
		filter *Filter,
		wait *Wait,
		options *GetSingleEntryOptions,
	) (*Data, error)
	GetJsonData(
		ctx context.Context,
		projection string,
		authData *AuthData,
		id string,
		filter *JsonFilter,
		wait *JsonWait,
		options *GetSingleEntryOptions,
	) (*JsonData, error)
	GetViewData(
		ctx context.Context,
		view string,
		authData *AuthData,
		filter *Filter,
		wait *Wait,
		options *GetSingleEntryOptions,
	) (*Data, error)
	GetViewJsonData(
		ctx context.Context,
		view string,
		authData *AuthData,
		filter *JsonFilter,
		wait *JsonWait,
		options *GetSingleEntryOptions,
	) (*JsonData, error)
	GetDataList(
		ctx context.Context,
		projection string,
		authData *AuthData,
		pagination *Pagination,
		filter *Filter,
		order []Order,
		wait *ListWait,
		options *GetEntryOptions,
	) (*DataList, error)
	GetJsonDataList(
		ctx context.Context,
		projection string,
		authData *AuthData,
		pagination *Pagination,
		filter *JsonFilter,
		order []Order,
		wait *ListWait,
		options *GetEntryOptions,
	) (*JsonDataList, error)
	GetViewDataList(
		ctx context.Context,
		view string,
		authData *AuthData,
		pagination *Pagination,
		filter *Filter,
		order []Order,
		wait *ListWait,
		options *GetEntryOptions,
	) (*DataList, error)
	GetViewJsonDataList(
		ctx context.Context,
		view string,
		authData *AuthData,
		pagination *Pagination,
		filter *JsonFilter,
		order []Order,
		wait *ListWait,
		options *GetEntryOptions,
	) (*JsonDataList, error)
	UpsertData(
		ctx context.Context,
		projection string,
		authData *AuthData,
		id string,
		payload Data,
		eventMetadata *EventMetadata,
	) (*UpsertResponse, error)
	UpsertJsonData(
		ctx context.Context,
		projection string,
		authData *AuthData,
		id string,
		payload JsonData,
		eventMetadata *EventMetadata,
	) (*JsonUpsertResponse, error)
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
	DeleteDataByJsonFilter(
		ctx context.Context,
		projection string,
		authData *AuthData,
		filter *JsonFilter,
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
	wait *Wait,
	options *GetSingleEntryOptions,
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
		Wait:                      wait.toDeliveryWait(),
		Target:                    options.Target(),
		ReturnEmptyDataIfNotFound: options.ReturnEmptyDataIfNotFound(),
		UseStrongConsistency:      options.UseStrongConsistency(),
		UseStrongConsistencyById:  options.UseStrongConsistencyById(),
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

func (c *projectionsDeliveryClient) GetJsonData(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	filter *JsonFilter,
	wait *JsonWait,
	options *GetSingleEntryOptions,
) (*JsonData, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetData(ctx, deliverypb.GetDataRequest_builder{
		Projection:                projection,
		Auth:                      pbAuthData,
		DataId:                    id,
		Filter:                    filter.toProtobufFilter(),
		Wait:                      wait.toDeliveryWait(),
		Target:                    options.Target(),
		ReturnEmptyDataIfNotFound: options.ReturnEmptyDataIfNotFound(),
		UseStrongConsistency:      options.UseStrongConsistency(),
		UseStrongConsistencyById:  options.UseStrongConsistencyById(),
	}.Build())
	if err != nil {
		return nil, err
	}

	result := response.GetResult()

	if result == nil {
		return nil, nil
	}

	data := JsonData(result.GetData())
	return &data, nil
}

func (c *projectionsDeliveryClient) GetViewData(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *Filter,
	wait *Wait,
	options *GetSingleEntryOptions,
) (*Data, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetViewData(ctx, deliverypb.GetViewDataRequest_builder{
		View:                      view,
		Auth:                      pbAuthData,
		Filter:                    filter.toProtobufFilter(),
		Wait:                      wait.toDeliveryWait(),
		Target:                    options.Target(),
		UseStrongConsistency:      options.UseStrongConsistency(),
		UseStrongConsistencyById:  options.UseStrongConsistencyById(),
		ReturnEmptyDataIfNotFound: options.ReturnEmptyDataIfNotFound(),
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

func (c *projectionsDeliveryClient) GetViewJsonData(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *JsonFilter,
	wait *JsonWait,
	options *GetSingleEntryOptions,
) (*JsonData, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetViewData(ctx, deliverypb.GetViewDataRequest_builder{
		View:                      view,
		Auth:                      pbAuthData,
		Filter:                    filter.toProtobufFilter(),
		Wait:                      wait.toDeliveryWait(),
		Target:                    options.Target(),
		UseStrongConsistency:      options.UseStrongConsistency(),
		UseStrongConsistencyById:  options.UseStrongConsistencyById(),
		ReturnEmptyDataIfNotFound: options.ReturnEmptyDataIfNotFound(),
	}.Build())
	if err != nil {
		return nil, err
	}

	result := response.GetResult()

	if result == nil {
		return nil, nil
	}
	data := JsonData(result.GetData())
	return &data, nil
}

func (c *projectionsDeliveryClient) GetDataList(
	ctx context.Context,
	projection string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	wait *ListWait,
	options *GetEntryOptions,
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
		Projection:               projection,
		Auth:                     pbAuthData,
		Limit:                    limit,
		Page:                     page,
		Filter:                   filter.toProtobufFilter(),
		Order:                    toProtobufOrder(order),
		Wait:                     wait.toDeliveryListWait(),
		Target:                   options.Target(),
		UseStrongConsistency:     options.UseStrongConsistency(),
		UseStrongConsistencyById: options.UseStrongConsistencyById(),
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

func (c *projectionsDeliveryClient) GetJsonDataList(
	ctx context.Context,
	projection string,
	authData *AuthData,
	pagination *Pagination,
	filter *JsonFilter,
	order []Order,
	wait *ListWait,
	options *GetEntryOptions,
) (*JsonDataList, error) {
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
		Projection:               projection,
		Auth:                     pbAuthData,
		Limit:                    limit,
		Page:                     page,
		Filter:                   filter.toProtobufFilter(),
		Order:                    toProtobufOrder(order),
		Wait:                     wait.toDeliveryListWait(),
		Target:                   options.Target(),
		UseStrongConsistency:     options.UseStrongConsistency(),
		UseStrongConsistencyById: options.UseStrongConsistencyById(),
	}.Build())
	if err != nil {
		return nil, err
	}

	var result []JsonData

	for _, element := range response.GetResult() {
		result = append(result, element.GetData())
	}

	return &JsonDataList{
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
	wait *ListWait,
	options *GetEntryOptions,
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
		View:                     view,
		Auth:                     pbAuthData,
		Limit:                    limit,
		Page:                     page,
		Filter:                   filter.toProtobufFilter(),
		Order:                    toProtobufOrder(order),
		Wait:                     wait.toDeliveryListWait(),
		Target:                   options.Target(),
		UseStrongConsistency:     options.UseStrongConsistency(),
		UseStrongConsistencyById: options.UseStrongConsistencyById(),
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

func (c *projectionsDeliveryClient) GetViewJsonDataList(
	ctx context.Context,
	view string,
	authData *AuthData,
	pagination *Pagination,
	filter *JsonFilter,
	order []Order,
	wait *ListWait,
	options *GetEntryOptions,
) (*JsonDataList, error) {
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
		View:                     view,
		Auth:                     pbAuthData,
		Limit:                    limit,
		Page:                     page,
		Filter:                   filter.toProtobufFilter(),
		Order:                    toProtobufOrder(order),
		Wait:                     wait.toDeliveryListWait(),
		Target:                   options.Target(),
		UseStrongConsistency:     options.UseStrongConsistency(),
		UseStrongConsistencyById: options.UseStrongConsistencyById(),
	}.Build())
	if err != nil {
		return nil, err
	}

	var result []JsonData

	for _, element := range response.GetResult() {
		result = append(result, element.GetData())
	}

	return &JsonDataList{
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

func (c *projectionsDeliveryClient) UpsertJsonData(
	ctx context.Context,
	projection string,
	authData *AuthData,
	id string,
	payload JsonData,
	eventMetadata *EventMetadata,
) (*JsonUpsertResponse, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.Upsert(ctx, deliverypb.UpsertRequest_builder{
		Projection:    projection,
		Auth:          pbAuthData,
		DataId:        id,
		Payload:       payload,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return nil, err
	}

	return &JsonUpsertResponse{
		ValidationErrors:      response.GetValidationErrors(),
		FieldValidationErrors: response.GetFieldValidationErrors(),
		Data:                  response.GetNewData(),
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

func (c *projectionsDeliveryClient) DeleteDataByJsonFilter(
	ctx context.Context,
	projection string,
	authData *AuthData,
	filter *JsonFilter,
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
