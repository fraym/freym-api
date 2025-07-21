package crud

import (
	"time"

	"github.com/fraym/freym-api/go/proto/crud/deliverypb"
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

type ListWait struct {
	Condition        *ListWaitCondition
	ConditionTimeout time.Duration
}

type ListWaitCondition struct {
	Operation string
	Value     int64
}

func (w *ListWait) toDeliveryListWait() *deliverypb.DataListWait {
	if w == nil || w.Condition == nil {
		return nil
	}

	return deliverypb.DataListWait_builder{
		Condition: deliverypb.DataListWaitCondition_builder{
			Operation: w.Condition.Operation,
			Value:     w.Condition.Value,
		}.Build(),
		Timeout: w.ConditionTimeout.Milliseconds(),
	}.Build()
}
