package crud

import (
	"github.com/fraym/freym-api/go/proto/crud/deliverypb"
	"github.com/samber/lo"
)

type Order struct {
	Field      string
	Descending bool
}

func toProtobufOrder(order []Order) []*deliverypb.DataOrder {
	if len(order) == 0 {
		return nil
	}

	return lo.Map(order, func(o Order, _ int) *deliverypb.DataOrder {
		return deliverypb.DataOrder_builder{
			Field:      o.Field,
			Descending: o.Descending,
		}.Build()
	})
}
