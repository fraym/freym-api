package crud

import (
	"context"
	"time"

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

type UpdateResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
	Entry                 Entry
}

type CloneResponse struct {
	ValidationErrors      []string
	FieldValidationErrors map[string]string
	Entry                 Entry
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
	GetEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		filter *Filter,
		wait *Wait,
		returnEmptyDataIfNotFound bool,
		useStrongConsistency bool,
		target deliverypb.DeploymentTarget,
	) (*Entry, error)
	GetEntryList(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		pagination *Pagination,
		filter *Filter,
		order []Order,
		useStrongConsistency bool,
		target deliverypb.DeploymentTarget,
	) (*EntryList, error)
	CreateEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		data Entry,
		eventMetadata *EventMetadata,
	) (*CreateResponse, error)
	UpdateEntry(
		ctx context.Context,
		typeName string,
		authData *AuthData,
		id string,
		data Entry,
		eventMetadata *EventMetadata,
	) (*UpdateResponse, error)
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
	returnEmptyDataIfNotFound bool,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
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
		ReturnEmptyDataIfNotFound: returnEmptyDataIfNotFound,
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
	return getEntryMap(result.GetData())
}

func (c *crudDeliveryClient) GetEntryList(
	ctx context.Context,
	typeName string,
	authData *AuthData,
	pagination *Pagination,
	filter *Filter,
	order []Order,
	useStrongConsistency bool,
	target deliverypb.DeploymentTarget,
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
		Type:                 typeName,
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
