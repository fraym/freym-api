package crud

import (
	"context"

	"github.com/Becklyn/go-wire-core/json"

	"github.com/fraym/freym-api/go/crud/config"
	"github.com/fraym/freym-api/go/crud/delivery"
	"github.com/fraym/freym-api/go/proto/crud/deliverypb"
)

type EntryList struct {
	Limit   int64
	Page    int64
	Total   int64
	Entries []Entry
}

type Entry map[string]any

type CreateResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
	Entry                 Entry
}

type UpsertResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
	Entry                 Entry
}

type UpdateResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
	Entry                 Entry
}

type UpdateByFilterResponse struct {
	NumberOfUpdatedEntries int64
	ValidationErrors       map[string]UpdateByFilterDataValidationErrorResponse
}

type UpdateByFilterDataValidationErrorResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
}

type CloneResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
	Entry                 Entry
}

type DeliveryClient interface {
	GetEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		filter *Filter,
		wait *Wait,
		options *GetSingleEntryOptions,
	) (*Entry, error)
	GetEntryList(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		pagination *Pagination,
		filter *Filter,
		order []Order,
		wait *ListWait,
		options *GetEntryOptions,
	) (*EntryList, error)
	GetViewEntry(
		ctx context.Context,
		view string,
		authData *AuthData,
		filter *Filter,
		wait *Wait,
		options *GetEntryOptions,
	) (*Entry, error)
	GetViewEntryList(
		ctx context.Context,
		view string,
		authData *AuthData,
		pagination *Pagination,
		filter *Filter,
		order []Order,
		wait *ListWait,
		options *GetEntryOptions,
	) (*EntryList, error)
	CreateEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		data Entry,
		eventMetadata *EventMetadata,
	) (*CreateResponse, error)
	UpsertEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		data Entry,
		eventMetadata *EventMetadata,
	) (*UpsertResponse, error)
	UpdateEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		data Entry,
		eventMetadata *EventMetadata,
	) (*UpdateResponse, error)
	UpdateEntryByFilter(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		filter *Filter,
		data Entry,
		eventMetadata *EventMetadata,
	) (*UpdateByFilterResponse, error)
	CloneEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		newId string,
		eventMetadata *EventMetadata,
	) (*CloneResponse, error)
	DeleteEntryById(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		eventMetadata *EventMetadata,
	) (int64, error)
	DeleteEntryByFilter(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		filter *Filter,
		eventMetadata *EventMetadata,
	) (int64, error)
	Close() error
}

type crudDeliveryClient struct {
	clientCloseFn func() error
	client        deliverypb.ServiceClient
}

func NewDeliveryClient(conf *config.Config) (DeliveryClient, error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	client, clientCloseFn, err := delivery.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &crudDeliveryClient{
		clientCloseFn: clientCloseFn,
		client:        client,
	}, nil
}

func (c *crudDeliveryClient) GetEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	filter *Filter,
	wait *Wait,
	options *GetSingleEntryOptions,
) (*Entry, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetData(ctx, deliverypb.GetDataRequest_builder{
		Type:                      typeName,
		Auth:                      pbAuthData,
		Id:                        id,
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
	return getEntryMap(result.GetData())
}

func (c *crudDeliveryClient) GetEntryList(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	wait *ListWait,
	options *GetEntryOptions,
) (*EntryList, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	var (
		limit int64
		page  int64
	)

	if pagination != nil {
		limit = pagination.Limit
		page = pagination.Page
	}

	response, err := c.client.GetDataList(ctx, deliverypb.GetDataListRequest_builder{
		Type:                     typeName,
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

	var result []Entry

	for _, element := range response.GetResult() {
		data, err := getEntryMap(element.GetData())
		if err != nil {
			return nil, err
		}

		result = append(result, *data)
	}

	return &EntryList{
		Limit:   response.GetLimit(),
		Page:    response.GetPage(),
		Total:   response.GetTotal(),
		Entries: result,
	}, nil
}

func (c *crudDeliveryClient) GetViewEntry(
	ctx context.Context,
	view string,
	authData *AuthData,
	filter *Filter,
	wait *Wait,
	options *GetEntryOptions,
) (*Entry, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.GetViewData(ctx, deliverypb.GetViewDataRequest_builder{
		View:                     view,
		Auth:                     pbAuthData,
		Filter:                   filter.toProtobufFilter(),
		Wait:                     wait.toDeliveryWait(),
		Target:                   options.Target(),
		UseStrongConsistency:     options.UseStrongConsistency(),
		UseStrongConsistencyById: options.UseStrongConsistencyById(),
	}.Build())
	if err != nil {
		return nil, err
	}

	result := response.GetResult()

	if result == nil {
		return nil, nil
	}
	return getEntryMap(result.GetData())
}

func (c *crudDeliveryClient) GetViewEntryList(
	ctx context.Context,
	view string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	wait *ListWait,
	options *GetEntryOptions,
) (*EntryList, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	var (
		limit int64
		page  int64
	)

	if pagination != nil {
		limit = pagination.Limit
		page = pagination.Page
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

	var result []Entry

	for _, element := range response.GetResult() {
		data, err := getEntryMap(element.GetData())
		if err != nil {
			return nil, err
		}

		result = append(result, *data)
	}

	return &EntryList{
		Limit:   response.GetLimit(),
		Page:    response.GetPage(),
		Total:   response.GetTotal(),
		Entries: result,
	}, nil
}

func (c *crudDeliveryClient) Close() error {
	return c.clientCloseFn()
}

func getEntryMap(data map[string]string) (*Entry, error) {
	entryMap := Entry{}

	for key, value := range data {
		var originalData any

		if err := json.Unmarshal([]byte(value), &originalData); err != nil {
			return nil, err
		}

		entryMap[key] = originalData
	}

	return &entryMap, nil
}

func getEntryStringMap(data Entry) (map[string]string, error) {
	dataMap := map[string]string{}

	for key, value := range data {
		bytes, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}

		dataMap[key] = string(bytes)
	}

	return dataMap, nil
}

func (c *crudDeliveryClient) CreateEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	data Entry,
	eventMetadata *EventMetadata,
) (*CreateResponse, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	entryMap, err := getEntryStringMap(data)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Create(ctx, deliverypb.CreateRequest_builder{
		Type:          typeName,
		Auth:          pbAuthData,
		Id:            id,
		Data:          entryMap,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return nil, err
	}

	entry, err := getEntryMap(response.GetNewData())
	if err != nil {
		return nil, err
	}

	return &CreateResponse{
		ValidationErrors:      response.GetValidationErrors(),
		FieldValidationErrors: response.GetFieldValidationErrors(),
		Entry:                 *entry,
	}, nil
}

func (c *crudDeliveryClient) UpsertEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	data Entry,
	eventMetadata *EventMetadata,
) (*UpsertResponse, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	entryMap, err := getEntryStringMap(data)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Upsert(ctx, deliverypb.UpsertRequest_builder{
		Type:          typeName,
		Auth:          pbAuthData,
		Id:            id,
		Data:          entryMap,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return nil, err
	}

	entry, err := getEntryMap(response.GetNewData())
	if err != nil {
		return nil, err
	}

	return &UpsertResponse{
		ValidationErrors:      response.GetValidationErrors(),
		FieldValidationErrors: response.GetFieldValidationErrors(),
		Entry:                 *entry,
	}, nil
}

func (c *crudDeliveryClient) UpdateEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	data Entry,
	eventMetadata *EventMetadata,
) (*UpdateResponse, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	entryMap, err := getEntryStringMap(data)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Update(ctx, deliverypb.UpdateRequest_builder{
		Type:          typeName,
		Auth:          pbAuthData,
		Id:            id,
		Data:          entryMap,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return nil, err
	}

	entry, err := getEntryMap(response.GetNewData())
	if err != nil {
		return nil, err
	}

	return &UpdateResponse{
		ValidationErrors:      response.GetValidationErrors(),
		FieldValidationErrors: response.GetFieldValidationErrors(),
		Entry:                 *entry,
	}, nil
}

func (c *crudDeliveryClient) UpdateEntryByFilter(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	filter *Filter,
	data Entry,
	eventMetadata *EventMetadata,
) (*UpdateByFilterResponse, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	entryMap, err := getEntryStringMap(data)
	if err != nil {
		return nil, err
	}

	response, err := c.client.UpdateByFilter(ctx, deliverypb.UpdateByFilterRequest_builder{
		Type:          typeName,
		Auth:          pbAuthData,
		Filter:        filter.toProtobufFilter(),
		Data:          entryMap,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return nil, err
	}

	var validationErrors map[string]UpdateByFilterDataValidationErrorResponse

	if response.GetValidationErrors() != nil {
		validationErrors = map[string]UpdateByFilterDataValidationErrorResponse{}

		for key, value := range response.GetValidationErrors() {
			validationErrors[key] = UpdateByFilterDataValidationErrorResponse{
				ValidationErrors:      value.GetValidationErrors(),
				FieldValidationErrors: value.GetFieldValidationErrors(),
			}
		}
	}

	return &UpdateByFilterResponse{
		NumberOfUpdatedEntries: response.GetNumberOfUpdatedEntries(),
		ValidationErrors:       validationErrors,
	}, nil
}

func (c *crudDeliveryClient) CloneEntry(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	newId string,
	eventMetadata *EventMetadata,
) (*CloneResponse, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return nil, err
	}

	response, err := c.client.Clone(ctx, deliverypb.CloneRequest_builder{
		Type:          typeName,
		Auth:          pbAuthData,
		Id:            id,
		NewId:         newId,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return nil, err
	}

	entry, err := getEntryMap(response.GetNewData())
	if err != nil {
		return nil, err
	}

	return &CloneResponse{
		ValidationErrors:      response.GetValidationErrors(),
		FieldValidationErrors: response.GetFieldValidationErrors(),
		Entry:                 *entry,
	}, nil
}

func (c *crudDeliveryClient) DeleteEntryById(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	id string,
	eventMetadata *EventMetadata,
) (int64, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return 0, err
	}

	response, err := c.client.Delete(ctx, deliverypb.DeleteRequest_builder{
		Type:          typeName,
		Auth:          pbAuthData,
		Id:            id,
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return 0, err
	}
	return response.GetNumberOfDeletedEntries(), nil
}

func (c *crudDeliveryClient) DeleteEntryByFilter(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	filter *Filter,
	eventMetadata *EventMetadata,
) (int64, error) {
	pbAuthData, err := authData.getProtobufAuthData()
	if err != nil {
		return 0, err
	}

	response, err := c.client.Delete(ctx, deliverypb.DeleteRequest_builder{
		Type:          typeName,
		Auth:          pbAuthData,
		Filter:        filter.toProtobufFilter(),
		EventMetadata: eventMetadata.getProtobufEventMetadata(),
	}.Build())
	if err != nil {
		return 0, err
	}
	return response.GetNumberOfDeletedEntries(), nil
}
