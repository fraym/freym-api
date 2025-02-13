package parser

import (
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

func GetValue(value ast.Value) (any, error) {
	switch value.GetKind() {
	case kinds.StringValue:
		return value.GetValue(), nil
	case kinds.EnumValue:
		return value.GetValue(), nil
	case kinds.BooleanValue:
		return value.GetValue(), nil
	case kinds.IntValue:
		return strconv.Atoi(value.GetValue().(string))
	case kinds.FloatValue:
		return strconv.ParseFloat(value.GetValue().(string), 64)
	case kinds.ListValue:
		var values []any

		for _, astValue := range value.GetValue().([]ast.Value) {
			value, err := GetValue(astValue)
			if err != nil {
				return nil, err
			}

			values = append(values, value)
		}

		return values, nil
	case kinds.ObjectValue:
		values := map[string]any{}

		for _, astObjectField := range value.GetValue().([]*ast.ObjectField) {
			value, err := GetValue(astObjectField.Value)
			if err != nil {
				return nil, err
			}

			values[astObjectField.Name.Value] = value
		}

		return values, nil
	default:
		return nil, fmt.Errorf("unknown value type `%s`", value.GetKind())
	}
}
