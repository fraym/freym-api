package util

import (
	"encoding/json"
	"time"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/freym-api/go/streams/domain/dto"
)

func SubscriptionEventsFromPb(data []*managementpb.Event) ([]*dto.SubscriptionEvent, error) {
	var events []*dto.SubscriptionEvent

	for _, eventData := range data {
		event, err := SubscriptionEventFromPb(eventData)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func SubscriptionEventFromPb(data *managementpb.Event) (*dto.SubscriptionEvent, error) {
	payload := map[string]dto.EventPayload{}

	for key, value := range data.GetPayload() {
		var valueData any
		err := json.Unmarshal([]byte(value.GetValue()), &valueData)
		if err != nil {
			return nil, err
		}

		var gdprPayload *dto.GdprEventPayload

		gdprData := value.GetGdpr()
		if gdprData != nil {
			var defaultValue any
			err := json.Unmarshal([]byte(gdprData.GetDefault()), &defaultValue)
			if err != nil {
				return nil, err
			}

			gdprPayload = &dto.GdprEventPayload{
				Id:          gdprData.GetId(),
				Default:     defaultValue,
				Invalidated: gdprData.GetIsInvalidated(),
			}
		}

		payload[key] = dto.EventPayload{
			Value: valueData,
			Gdpr:  gdprPayload,
		}
	}

	var (
		correlationId string
		causationId   string
		orderSerial   int64
		userId        string
		deploymentId  int64
		retryNumber   int64
	)

	if meta := data.GetMetadata(); meta != nil {
		correlationId = meta.GetCorrelationId()
		causationId = meta.GetCausationId()
		orderSerial = meta.GetOrderSerial()
		userId = meta.GetUserId()
		deploymentId = meta.GetDeploymentId()
		retryNumber = meta.GetRetryNumber()
	}

	return &dto.SubscriptionEvent{
		Id:            data.GetId(),
		TenantId:      data.GetTenantId(),
		Stream:        data.GetStream(),
		Type:          data.GetType(),
		CorrelationId: correlationId,
		CausationId:   causationId,
		Reason:        data.GetReason(),
		Payload:       payload,
		Topic:         data.GetTopic(),
		RaisedAt:      time.Unix(0, data.GetRaisedAt()),
		OrderSerial:   orderSerial,
		UserId:        userId,
		DeploymentId:  deploymentId,
		RetryNumber:   retryNumber,
	}, nil
}

func ErroneousEventsFromPb(data []*managementpb.ErroneousEvent) ([]*dto.ErroneousEvent, error) {
	var events []*dto.ErroneousEvent

	for _, eventData := range data {
		event, err := SubscriptionEventFromPb(eventData.GetEvent())
		if err != nil {
			return nil, err
		}

		events = append(events, &dto.ErroneousEvent{
			Event:         event,
			ConsumerGroup: eventData.GetConsumerGroup(),
			Error:         eventData.GetError(),
		})
	}

	return events, nil
}
