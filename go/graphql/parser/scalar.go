package parser

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/ast"
)

func GetScalarFromDefinition(definition *ast.ScalarDefinition) (*schema.Scalar, error) {
	if definition == nil || definition.Name == nil || len(definition.Name.Value) == 0 {
		return nil, fmt.Errorf("cannot parse enum without name")
	}

	directives, err := GetDirectivesFromScalarDefinition(definition)
	if err != nil {
		return nil, err
	}

	return &schema.Scalar{
		Name:       definition.Name.Value,
		Directives: directives,
	}, nil
}
