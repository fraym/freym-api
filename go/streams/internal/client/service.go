package client

import (
	"context"
	"strings"
	"time"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/freym-api/go/streams/config"
	"github.com/fraym/freym-api/go/streams/domain/dto"
	"github.com/fraym/freym-api/go/streams/internal/util"
	"github.com/fraym/golog"
)

type Service struct {
	ctx        context.Context
	config     *config.Config
	client     managementpb.ServiceClient
	logger     golog.Logger
	retryPause time.Duration
}

func NewService(
	ctx context.Context,
	config *config.Config,
	client managementpb.ServiceClient,
	logger golog.Logger,
) (*Service, error) {
	retryPause := time.Millisecond * 100

	return &Service{
		ctx:        ctx,
		config:     config,
		client:     client,
		logger:     logger,
		retryPause: retryPause,
	}, nil
}

func (s *Service) IterateAllEvents(
	ctx context.Context,
	tenant string,
	topic string,
	types []string,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	lastEventCheck, err := s.getLastEventCheckFunc(ctx, tenant, topic)
	if err != nil {
		return err
	}
	if lastEventCheck == nil {
		return nil
	}

	lastEventId := ""

	return util.Iterate(ctx, handler, func(ctx context.Context, page, perPage int) ([]*dto.SubscriptionEvent, error) {
		var events []*managementpb.Event

		if lastEventId == "" {
			response, err := util.RetryWithResult(func() (*managementpb.PaginateEventsResponse, error) {
				return s.client.PaginateEvents(ctx, managementpb.PaginateEventsRequest_builder{
					TenantId: tenant,
					Topic:    topic,
					Types:    types,
					Page:     0,
					PerPage:  int64(perPage),
				}.Build())
			}, s.retryPause, 50)
			if err != nil {
				return nil, err
			}

			events = response.GetEvents()
		} else {
			response, err := util.RetryWithResult(func() (*managementpb.PaginateEventsAfterEventIdResponse, error) {
				return s.client.PaginateEventsAfterEventId(ctx, managementpb.PaginateEventsAfterEventIdRequest_builder{
					TenantId: tenant,
					Topic:    topic,
					Types:    types,
					Page:     0,
					PerPage:  int64(perPage),
					EventId:  lastEventId,
				}.Build())
			}, s.retryPause, 50)
			if err != nil {
				return nil, err
			}

			events = response.GetEvents()
		}

		if len(events) > 0 {
			lastEventId = events[len(events)-1].GetId()
		}

		return util.SubscriptionEventsFromPb(events)
	}, lastEventCheck, perPage, queueSize)
}

func (s *Service) IterateAllEventsAfterEvent(
	ctx context.Context,
	tenant string,
	topic string,
	types []string,
	eventId string,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	lastEventCheck, err := s.getLastEventCheckFunc(ctx, tenant, topic)
	if err != nil {
		return err
	}
	if lastEventCheck == nil {
		return nil
	}

	return util.Iterate(ctx, handler, func(ctx context.Context, page, perPage int) ([]*dto.SubscriptionEvent, error) {
		response, err := util.RetryWithResult(func() (*managementpb.PaginateEventsAfterEventIdResponse, error) {
			return s.client.PaginateEventsAfterEventId(ctx, managementpb.PaginateEventsAfterEventIdRequest_builder{
				TenantId: tenant,
				Topic:    topic,
				Types:    types,
				EventId:  eventId,
				Page:     int64(page),
				PerPage:  int64(perPage),
			}.Build())
		}, s.retryPause, 50)
		if err != nil {
			return nil, err
		}

		return util.SubscriptionEventsFromPb(response.GetEvents())
	}, lastEventCheck, perPage, queueSize)
}

func (s *Service) IterateStream(
	ctx context.Context,
	tenant string,
	topic string,
	stream string,
	deploymentId int64,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	lastEventCheck, err := s.getLastEventCheckFunc(ctx, tenant, topic)
	if err != nil {
		return err
	}
	if lastEventCheck == nil {
		return nil
	}

	snapshotEventId := ""

	return util.Iterate(ctx, handler, func(ctx context.Context, page, perPage int) ([]*dto.SubscriptionEvent, error) {
		response, err := util.RetryWithResult(func() (*managementpb.PaginateStreamResponse, error) {
			return s.client.PaginateStream(ctx, managementpb.PaginateStreamRequest_builder{
				TenantId:        tenant,
				Topic:           topic,
				Stream:          stream,
				Page:            int64(page),
				PerPage:         int64(perPage),
				DeploymentId:    deploymentId,
				SnapshotEventId: snapshotEventId,
			}.Build())
		}, s.retryPause, 50)
		if err != nil {
			return nil, err
		}

		snapshotEventId = response.GetSnapshotEventId()

		return util.SubscriptionEventsFromPb(response.GetEvents())
	}, lastEventCheck, perPage, queueSize)
}

func (s *Service) IterateStreamAfterEvent(
	ctx context.Context,
	tenant string,
	topic string,
	stream string,
	eventId string,
	deploymentId int64,
	perPage int,
	queueSize int,
	handler dto.HandlerFunc,
) error {
	lastEventCheck, err := s.getLastEventCheckFunc(ctx, tenant, topic)
	if err != nil {
		return err
	}
	if lastEventCheck == nil {
		return nil
	}

	snapshotEventId := ""

	return util.Iterate(ctx, handler, func(ctx context.Context, page, perPage int) ([]*dto.SubscriptionEvent, error) {
		response, err := util.RetryWithResult(func() (*managementpb.PaginateStreamAfterEventIdResponse, error) {
			return s.client.PaginateStreamAfterEventId(ctx, managementpb.PaginateStreamAfterEventIdRequest_builder{
				TenantId:        tenant,
				Topic:           topic,
				Stream:          stream,
				EventId:         eventId,
				Page:            int64(page),
				PerPage:         int64(perPage),
				DeploymentId:    deploymentId,
				SnapshotEventId: snapshotEventId,
			}.Build())
		}, s.retryPause, 50)
		if err != nil {
			return nil, err
		}

		snapshotEventId = response.GetSnapshotEventId()

		return util.SubscriptionEventsFromPb(response.GetEvents())
	}, lastEventCheck, perPage, queueSize)
}

func (s *Service) IsStreamEmpty(ctx context.Context, tenant string, topic string, stream string) (bool, error) {
	response, err := util.RetryWithResult(func() (*managementpb.IsStreamEmptyResponse, error) {
		return s.client.IsStreamEmpty(ctx, managementpb.IsStreamEmptyRequest_builder{
			TenantId: tenant,
			Topic:    topic,
			Stream:   stream,
		}.Build())
	}, s.retryPause, 50)
	if err != nil {
		return false, err
	}

	return response.GetIsEmpty(), nil
}

func (s *Service) GetEvent(
	ctx context.Context,
	tenantId string,
	topic string,
	eventId string,
) (*dto.SubscriptionEvent, error) {
	response, err := util.RetryWithResult(func() (*managementpb.Event, error) {
		return s.client.GetEvent(ctx, managementpb.GetEventRequest_builder{
			TenantId: tenantId,
			Topic:    topic,
			Id:       eventId,
		}.Build())
	}, s.retryPause, 50)
	if err != nil {
		return nil, err
	}

	return util.SubscriptionEventFromPb(response)
}

func (s *Service) GetLastEvent(ctx context.Context, tenantId string, topic string) (*dto.SubscriptionEvent, error) {
	response, err := util.RetryWithResult(func() (*managementpb.Event, error) {
		res, err := s.client.GetLastEvent(ctx, managementpb.GetLastEventRequest_builder{
			TenantId: tenantId,
			Topic:    topic,
		}.Build())
		if err != nil {
			if isLastEventNotFoundError(err) {
				return nil, nil
			}
			return nil, err
		}
		return res, nil
	}, s.retryPause, 50)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, nil
	}

	return util.SubscriptionEventFromPb(response)
}

func (s *Service) GetLastEventByTypes(
	ctx context.Context,
	tenantId string,
	topic string,
	types []string,
) (*dto.SubscriptionEvent, error) {
	response, err := util.RetryWithResult(func() (*managementpb.Event, error) {
		res, err := s.client.GetLastEventByTypes(ctx, managementpb.GetLastEventByTypesRequest_builder{
			TenantId: tenantId,
			Topic:    topic,
			Types:    types,
		}.Build())
		if err != nil {
			if isLastEventNotFoundError(err) {
				return nil, nil
			}
			return nil, err
		}

		return res, nil
	}, s.retryPause, 50)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, nil
	}

	return util.SubscriptionEventFromPb(response)
}

func (s *Service) NewSubscription(topics []string, ignoreUnhandledEvents bool, deploymentId int64) *Subscription {
	return newSubscription(s.ctx, s.config, s.client, s.logger, topics, ignoreUnhandledEvents, deploymentId)
}

func (s *Service) Publish(ctx context.Context, topic string, events []*dto.PublishEvent) error {
	var eventsToPublish []*managementpb.PublishEvent

	for _, event := range events {
		publishEvent, err := event.ToProtobufPublishEvent()
		if err != nil {
			return err
		}
		eventsToPublish = append(eventsToPublish, publishEvent)
	}

	return util.Retry(func() error {
		_, err := s.client.Publish(ctx, managementpb.PublishRequest_builder{
			Topic:  topic,
			Events: eventsToPublish,
		}.Build())
		return err
	}, s.retryPause, 50)
}

func (s *Service) InvalidateGdprData(ctx context.Context, tenantId string, topic string, gdprId string) error {
	return util.Retry(func() error {
		_, err := s.client.InvalidateGdpr(ctx, managementpb.InvalidateGdprRequest_builder{
			TenantId: tenantId,
			Topic:    topic,
			GdprId:   gdprId,
		}.Build())
		return err
	}, s.retryPause, 50)
}

func (s *Service) IntroduceGdprOnEventField(
	ctx context.Context,
	tenantId string,
	topic string,
	eventId string,
	fieldName string,
	defaultValue string,
) error {
	return util.Retry(func() error {
		_, err := s.client.IntroduceGdprOnEventField(ctx, managementpb.IntroduceGdprOnEventFieldRequest_builder{
			TenantId:     tenantId,
			Topic:        topic,
			EventId:      eventId,
			FieldName:    fieldName,
			DefaultValue: defaultValue,
		}.Build())
		return err
	}, s.retryPause, 50)
}

func (s *Service) CreateStreamSnapshot(
	ctx context.Context,
	tenantId string,
	topic string,
	stream string,
	lastSnapshottedEventId string,
	snapshotEvent *dto.PublishEvent,
) error {
	publishEvent, err := snapshotEvent.ToProtobufPublishEvent()
	if err != nil {
		return err
	}

	return util.Retry(func() error {
		_, err := s.client.CreateStreamSnapshot(ctx, managementpb.CreateStreamSnapshotRequest_builder{
			TenantId:               tenantId,
			Topic:                  topic,
			Stream:                 stream,
			LastSnapshottedEventId: lastSnapshottedEventId,
			SnapshotEvent:          publishEvent,
		}.Build())
		return err
	}, s.retryPause, 50)
}

func (s *Service) RenameEventType(ctx context.Context, topic string, oldEventType string, newEventType string) error {
	return util.Retry(func() error {
		_, err := s.client.RenameEventType(ctx, managementpb.RenameEventTypeRequest_builder{
			Topic:   topic,
			OldType: oldEventType,
			NewType: newEventType,
		}.Build())
		return err
	}, s.retryPause, 50)
}

type LastEventCheckFunc func(ctx context.Context, lastEvent *dto.SubscriptionEvent) bool

func isLastEventNotFoundError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "unable to find last event")
}

func (s *Service) getLastEventCheckFunc(ctx context.Context, tenant string, topic string) (LastEventCheckFunc, error) {
	eventResponse, err := util.RetryWithResult(func() (*managementpb.Event, error) {
		res, err := s.client.GetLastEvent(ctx, managementpb.GetLastEventRequest_builder{
			TenantId: tenant,
			Topic:    topic,
		}.Build())
		if err != nil {
			if isLastEventNotFoundError(err) {
				return nil, nil
			}
			return nil, err
		}

		return res, nil
	}, s.retryPause, 50)
	if err != nil {
		return nil, err
	}
	if eventResponse == nil {
		return nil, nil
	}

	now := time.Now().Add(3 * time.Second)
	var lastOrderSerial int64 = -1
	metadata := eventResponse.GetMetadata()

	if metadata != nil {
		lastOrderSerial = metadata.GetOrderSerial()
	}

	return func(ctx context.Context, lastEvent *dto.SubscriptionEvent) bool {
		if lastEvent == nil {
			return true
		}

		if lastOrderSerial == -1 {
			return lastEvent.RaisedAt.After(now)
		}

		return lastEvent.OrderSerial > lastOrderSerial
	}, nil
}
