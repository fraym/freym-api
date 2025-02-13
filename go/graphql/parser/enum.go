package parser

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/ast"
)

func GetEnumFromDefinition(definition *ast.EnumDefinition) (*schema.Enum, error) {
	if definition == nil || definition.Name == nil || len(definition.Name.Value) == 0 {
		return nil, fmt.Errorf("cannot parse enum without name")
	}

	var values []string

	for _, value := range definition.Values {
		values = append(values, value.Name.Value)
	}

	return &schema.Enum{
		Name:   definition.Name.Value,
		Values: values,
	}, nil
}
