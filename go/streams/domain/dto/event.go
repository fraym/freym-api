package dto

import (
	"context"
	"fmt"
	"time"

	"github.com/Becklyn/go-wire-core/json"
	"github.com/fraym/freym-api/go/proto/streams/managementpb"
)

type GdprEventPayload struct {
	Id          string
	Default     any
	Invalidated bool
}

type EventPayload struct {
	Value any
	Gdpr  *GdprEventPayload
}

type EventPayloadMap map[string]EventPayload

type PublishGdprEventPayload struct {
	Id      string
	Default any
}

type PublishEventPayload struct {
	Value any
	Gdpr  *PublishGdprEventPayload
}

type PublishEventPayloadMap map[string]PublishEventPayload

func (m EventPayloadMap) Get(key string) (any, error) {
	if value, ok := m[key]; ok {
		return value.Value, nil
	}
	return nil, fmt.Errorf("key %s does not exist in event", key)
}

func GetPayloadValue[T any](m EventPayloadMap, key string) (T, error) {
	value, err := m.Get(key)
	if err != nil {
		var value T
		return value, err
	}

	if value == nil {
		var nilValue T
		return nilValue, nil
	}

	returnValue, ok := value.(T)
	if !ok {
		var errValue T
		return errValue, fmt.Errorf("key %s is not of type %T in event, got %T", key, errValue, value)
	}
	return returnValue, nil
}

type SubscriptionEvent struct {
	Id            string
	TenantId      string
	Stream        string
	Type          string
	CorrelationId string
	CausationId   string
	Reason        string
	Payload       EventPayloadMap
	Topic         string
	RaisedAt      time.Time
	OrderSerial   int64
	UserId        string
	DeploymentId  int64
}

type HandlerFunc = func(ctx context.Context, event *SubscriptionEvent) error

type PublishEvent struct {
	Id            string
	TenantId      string
	Stream        string
	Type          string
	CorrelationId string
	CausationId   string
	UserId        string
	DeploymentId  int64
	Reason        string
	Payload       PublishEventPayloadMap
	Broadcast     bool
}

func (e *PublishEvent) ToProtobufPublishEvent() (*managementpb.PublishEvent, error) {
	payload := map[string]*managementpb.PublishEventPayload{}

	for key, value := range e.Payload {
		valueString, err := json.Marshal(value.Value)
		if err != nil {
			return nil, err
		}

		var gdpr *managementpb.PublishEventGdprValue

		if value.Gdpr != nil {
			var defaultString string

			if value.Gdpr.Default != nil {
				newDefaultString, err := json.Marshal(value.Gdpr.Default)
				if err != nil {
					return nil, err
				}

				defaultString = string(newDefaultString)
			}

			gdpr = managementpb.PublishEventGdprValue_builder{
				Id:      value.Gdpr.Id,
				Default: defaultString,
			}.Build()
		}

		payload[key] = managementpb.PublishEventPayload_builder{
			Value: string(valueString),
			Gdpr:  gdpr,
		}.Build()
	}

	return managementpb.PublishEvent_builder{
		TenantId: e.TenantId,
		Id:       e.Id,
		Type:     e.Type,
		Stream:   e.Stream,
		Reason:   e.Reason,
		Options: managementpb.EventOptions_builder{
			Broadcast: e.Broadcast,
		}.Build(),
		Metadata: managementpb.EventMetadata_builder{
			CorrelationId: e.CorrelationId,
			CausationId:   e.CausationId,
			UserId:        e.UserId,
			DeploymentId:  e.DeploymentId,
		}.Build(),
		Payload: payload,
	}.Build(), nil
}
