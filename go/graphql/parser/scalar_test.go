package parser_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/stretchr/testify/assert"
)

func TestGetScalarsFormDefinition(t *testing.T) {
	expectedScalarName := "scalar"
	expectedDirectiveName := "directive"
	expectedArgName := "arg"
	expectedArgValue := "value"

	scalar, err := parser.GetScalarFromDefinition(ast.NewScalarDefinition(&ast.ScalarDefinition{
		Name: ast.NewName(&ast.Name{
			Value: expectedScalarName,
		}),
		Directives: []*ast.Directive{
			ast.NewDirective(&ast.Directive{
				Name: ast.NewName(&ast.Name{
					Value: expectedDirectiveName,
				}),
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
			}),
		},
	}))

	assert.NoError(t, err)
	assert.Equal(t, expectedScalarName, scalar.Name)
	assert.Equal(t, expectedDirectiveName, scalar.Directives[0].Name)
	assert.Equal(t, expectedArgName, scalar.Directives[0].Arguments[0].Name)
	assert.Equal(t, expectedArgValue, scalar.Directives[0].Arguments[0].Value)
}
