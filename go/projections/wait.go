package projections

import (
	"time"

	"github.com/fraym/freym-api/go/proto/projections/deliverypb"
)

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

type JsonWait struct {
	ConditionFilter  *JsonFilter
	ConditionTimeout time.Duration
}

func (w *JsonWait) toDeliveryWait() *deliverypb.DataWait {
	if w == nil || w.ConditionFilter == nil {
		return nil
	}

	return deliverypb.DataWait_builder{
		ConditionFilter: w.ConditionFilter.toProtobufFilter(),
		Timeout:         w.ConditionTimeout.Milliseconds(),
	}.Build()
}
