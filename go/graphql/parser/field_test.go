package parser_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func TestGetFieldsFormDefinition(t *testing.T) {
	expectedFieldName := "field"
	expectedFieldTypeName := "fieldType"
	expectedDirectiveName := "directive"
	expectedArgName := "arg"
	expectedArgValue := "value"

	fields, err := parser.GetFieldsFormDefinition(ast.NewObjectDefinition(&ast.ObjectDefinition{
		Fields: []*ast.FieldDefinition{
			ast.NewFieldDefinition(&ast.FieldDefinition{
				Name: ast.NewName(&ast.Name{
					Value: expectedFieldName,
				}),
				Type: ast.NewNamed(&ast.Named{
					Name: ast.NewName(&ast.Name{
						Value: expectedFieldTypeName,
					}),
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
			}),
		},
	}))

	assert.NoError(t, err)
	assert.Equal(t, expectedFieldName, fields[0].Name)
	assert.Equal(t, types.Comparable{kinds.Named, expectedFieldTypeName}, fields[0].Type)
	assert.Equal(t, expectedDirectiveName, fields[0].Directives[0].Name)
	assert.Equal(t, expectedArgName, fields[0].Directives[0].Arguments[0].Name)
	assert.Equal(t, expectedArgValue, fields[0].Directives[0].Arguments[0].Value)
}
