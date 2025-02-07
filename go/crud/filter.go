package crud

import (
	"github.com/Becklyn/go-wire-core/json"
	"github.com/fraym/freym-api/go/proto/crud/deliverypb"
)

type Filter struct {
	// all these field filters have to match
	Fields map[string]FieldFilter
	// these filters will be calculated using the and operator
	And []Filter
	// these filters will be calculated using the or operator
	Or []Filter
}

type FieldFilter struct {
	Operation string
	Type      string
	Value     any
}

func (f *Filter) toProtobufFilter() *deliverypb.DataFilter {
	if f == nil || (len(f.Fields) == 0 && len(f.And) == 0 && len(f.Or) == 0) {
		return nil
	}

	fields := map[string]*deliverypb.DataFieldFilter{}
	var and []*deliverypb.DataFilter
	var or []*deliverypb.DataFilter

	for fieldName, field := range f.Fields {
		fields[fieldName] = getFieldFilter(&field)
	}

	return deliverypb.DataFilter_builder{
		Fields: fields,
		And:    and,
		Or:     or,
	}.Build()
}

func getFieldFilter(field *FieldFilter) *deliverypb.DataFieldFilter {
	bytes, _ := json.Marshal(field.Value)

	return deliverypb.DataFieldFilter_builder{
		Operation: field.Operation,
		Type:      field.Type,
		Value:     string(bytes),
	}.Build()
}
