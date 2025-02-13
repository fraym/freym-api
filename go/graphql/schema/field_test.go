package schema_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func TestField_GetString(t *testing.T) {
	f := schema.Field{
		Name: "field",
		Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
		Directives: []schema.Directive{
			{
				Name:      "directive",
				Arguments: []schema.Argument{},
			},
			{
				Name:      "directive2",
				Arguments: []schema.Argument{},
			},
		},
	}

	expectedString := "\tfield: String! @directive @directive2\n"

	schemaString, err := f.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, schemaString)
}

func TestField_Equals(t *testing.T) {
	f := schema.Field{
		Name: "field",
		Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
		Directives: []schema.Directive{
			{
				Name: "directive2",
				Arguments: []schema.Argument{
					{
						Name:  "a",
						Value: "test",
					},
				},
			},
			{
				Name:      "directive",
				Arguments: []schema.Argument{},
			},
			{
				Name:      "directive2",
				Arguments: []schema.Argument{},
			},
		},
	}

	assert.True(t, f.Equals(&schema.Field{
		Name: "field",
		Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
		Directives: []schema.Directive{
			{
				Name:      "directive2",
				Arguments: []schema.Argument{},
			},
			{
				Name: "directive2",
				Arguments: []schema.Argument{
					{
						Name:  "a",
						Value: "test",
					},
				},
			},
			{
				Name:      "directive",
				Arguments: []schema.Argument{},
			},
		},
	}))

	assert.False(t, f.Equals(&schema.Field{
		Name: "field",
		Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
		Directives: []schema.Directive{
			{
				Name:      "directive2",
				Arguments: []schema.Argument{},
			},
			{
				Name: "directive2",
				Arguments: []schema.Argument{
					{
						Name:  "a",
						Value: "test2",
					},
				},
			},
			{
				Name:      "directive",
				Arguments: []schema.Argument{},
			},
		},
	}))
}
