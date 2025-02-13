package parser_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/stretchr/testify/assert"
)

func TestGetDirectivesFromObjectDefinition(t *testing.T) {
	expectedDirectiveName := "directive"
	expectedArgName := "arg"
	expectedArgValue := "value"

	directives, err := parser.GetDirectivesFromObjectDefinition(ast.NewObjectDefinition(&ast.ObjectDefinition{
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
	assert.Equal(t, expectedDirectiveName, directives[0].Name)
	assert.Equal(t, expectedArgName, directives[0].Arguments[0].Name)
	assert.Equal(t, expectedArgValue, directives[0].Arguments[0].Value)
}

func TestGetDirectivesFromFieldDefinition(t *testing.T) {
	expectedDirectiveName := "directive"
	expectedArgName := "arg"
	expectedArgValue := "value"

	directives, err := parser.GetDirectivesFromFieldDefinition(ast.NewFieldDefinition(&ast.FieldDefinition{
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
	assert.Equal(t, expectedDirectiveName, directives[0].Name)
	assert.Equal(t, expectedArgName, directives[0].Arguments[0].Name)
	assert.Equal(t, expectedArgValue, directives[0].Arguments[0].Value)
}

func TestGetDirectivesFromScalarDefinition(t *testing.T) {
	expectedDirectiveName := "directive"
	expectedArgName := "arg"
	expectedArgValue := "value"

	directives, err := parser.GetDirectivesFromScalarDefinition(ast.NewScalarDefinition(&ast.ScalarDefinition{
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
	assert.Equal(t, expectedDirectiveName, directives[0].Name)
	assert.Equal(t, expectedArgName, directives[0].Arguments[0].Name)
	assert.Equal(t, expectedArgValue, directives[0].Arguments[0].Value)
}
