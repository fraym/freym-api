package parser_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/stretchr/testify/assert"
)

func TestGetArgumentsOfDirective(t *testing.T) {
	expectedArgName := "arg"
	expectedArgValue := "value"

	args, err := parser.GetArgumentsOfDirective(ast.NewDirective(&ast.Directive{
		Arguments: []*ast.Argument{
			ast.NewArgument(&ast.Argument{
				Name: ast.NewName(&ast.Name{
					Value: expectedArgName,
				}),
				Value: ast.NewStringValue(&ast.StringValue{
					Value: expectedArgValue,
				}),
			}),
		},
	}))

	assert.NoError(t, err)
	assert.Equal(t, expectedArgName, args[0].Name)
	assert.Equal(t, expectedArgValue, args[0].Value)
}
