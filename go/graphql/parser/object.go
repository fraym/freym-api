package parser

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/ast"
)

func GetObjectFromDefinition(definition *ast.ObjectDefinition) (*schema.Object, error) {
	if definition == nil || definition.Name == nil || len(definition.Name.Value) == 0 {
		return nil, fmt.Errorf("cannot parse object without name")
	}

	directives, err := GetDirectivesFromObjectDefinition(definition)
	if err != nil {
		return nil, err
	}

	fields, err := GetFieldsFormDefinition(definition)
	if err != nil {
		return nil, err
	}

	return &schema.Object{
		Name:       definition.Name.Value,
		Directives: directives,
		Fields:     fields,
	}, nil
}
