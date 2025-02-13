package parser_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func TestGetObjectFromDefinition(t *testing.T) {
	expectedObjectName := "objectName"
	expectedDirectiveName := "directiveName"
	expectedFieldName := "fieldName"
	expectedFieldTypeName := "fieldTypeName"

	obj, err := parser.GetObjectFromDefinition(ast.NewObjectDefinition(&ast.ObjectDefinition{
		Name: ast.NewName(&ast.Name{Value: expectedObjectName}),
		Directives: []*ast.Directive{
			ast.NewDirective(&ast.Directive{Name: ast.NewName(&ast.Name{Value: expectedDirectiveName})}),
		},
		Fields: []*ast.FieldDefinition{
			ast.NewFieldDefinition(&ast.FieldDefinition{
				Name: ast.NewName(&ast.Name{Value: expectedFieldName}),
				Type: ast.NewNamed(&ast.Named{Name: ast.NewName(&ast.Name{Value: expectedFieldTypeName})}),
			}),
		},
	}))
	assert.NoError(t, err)
	assert.Equal(t, &schema.Object{
		Name: expectedObjectName,
		Directives: []schema.Directive{
			{Name: expectedDirectiveName},
		},
		Fields: []schema.Field{
			{
				Name: expectedFieldName,
				Type: types.Comparable{kinds.Named, expectedFieldTypeName},
			},
		},
	}, obj)
}
