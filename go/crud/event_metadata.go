package crud

import (
	"github.com/fraym/freym-api/go/proto/crud/deliverypb"
)

type EventMetadata struct {
	CorrelationId string
	CausationId   string
	UserId        string
	DeploymentId  int64
}

func (m *EventMetadata) getProtobufEventMetadata() *deliverypb.EventMetadata {
	if m == nil {
		return nil
	}

	return deliverypb.EventMetadata_builder{
		CorrelationId: m.CorrelationId,
		CausationId:   m.CausationId,
		DeploymentId:  m.DeploymentId,
		UserId:        m.UserId,
	}.Build()
}
