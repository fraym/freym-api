package projections

import (
	"time"

	"github.com/fraym/freym-api/go/proto/projections/deliverypb"
	"github.com/samber/lo"
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

type ListWait struct {
	Condition        []ListWaitCondition
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
		Condition: lo.Map(w.Condition, func(condition ListWaitCondition, _ int) *deliverypb.DataListWaitCondition {
			return deliverypb.DataListWaitCondition_builder{
				Operation: condition.Operation,
				Value:     condition.Value,
			}.Build()
		}),
		Timeout: w.ConditionTimeout.Milliseconds(),
	}.Build()
}
