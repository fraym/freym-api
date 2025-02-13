package parser

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/graphql-go/graphql/language/ast"
)

func GetFieldsFormDefinition(definition *ast.ObjectDefinition) ([]schema.Field, error) {
	var fields []schema.Field

	for _, field := range definition.Fields {
		if field == nil || field.Name == nil || len(field.Name.Value) == 0 {
			return nil, fmt.Errorf("cannot parse field without name on object `%s`", definition.Name.Value)
		}

		comparableType, err := types.GetComparable(field.Type)
		if err != nil {
			return nil, err
		}

		directives, err := GetDirectivesFromFieldDefinition(field)
		if err != nil {
			return nil, err
		}

		fields = append(fields, schema.Field{
			Name:       field.Name.Value,
			Type:       comparableType,
			Directives: directives,
		})
	}

	return fields, nil
}
