package projections

import (
	"github.com/fraym/freym-api/go/proto/projections/deliverypb"
)

type EventMetadata struct {
	CorrelationId string
	CausationId   string
	Target        deliverypb.DeploymentTarget
}

func (m *EventMetadata) getProtobufEventMetadata() *deliverypb.EventMetadata {
	if m == nil {
		return nil
	}

	return deliverypb.EventMetadata_builder{
		CorrelationId: m.CorrelationId,
		CausationId:   m.CausationId,
		Target:        m.Target,
	}.Build()
}
