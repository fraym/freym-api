package parser_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/stretchr/testify/assert"
)

func TestGetEnumFromDefinition(t *testing.T) {
	expectedEnumName := "enumName"
	expectedEnumValue := "value"

	obj, err := parser.GetEnumFromDefinition(ast.NewEnumDefinition(&ast.EnumDefinition{
		Name: ast.NewName(&ast.Name{Value: expectedEnumName}),
		Values: []*ast.EnumValueDefinition{
			ast.NewEnumValueDefinition(&ast.EnumValueDefinition{
				Name: ast.NewName(&ast.Name{Value: expectedEnumValue}),
			}),
		},
	}))
	assert.NoError(t, err)
	assert.Equal(t, &schema.Enum{
		Name:   expectedEnumName,
		Values: []string{expectedEnumValue},
	}, obj)
}
