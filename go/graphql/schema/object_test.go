package schema_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func TestObject_GetString(t *testing.T) {
	o := schema.Object{
		Name: "Object",
		Directives: []schema.Directive{
			{
				Name: "directive",
				Arguments: []schema.Argument{
					{
						Name:  "arg",
						Value: "test",
					},
				},
			},
		},
		Fields: []schema.Field{
			{
				Name: "field",
				Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
			},
		},
	}

	expectedString := "type Object @directive(arg: \"test\") {\n\tfield: String!\n}"

	schemaString, err := o.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, schemaString)
}

func TestObject_Equals(t *testing.T) {
	o := schema.Object{
		Name: "Object",
		Directives: []schema.Directive{
			{
				Name: "directive",
				Arguments: []schema.Argument{
					{
						Name:  "arg2",
						Value: "test",
					},
				},
			},
			{
				Name: "directive",
				Arguments: []schema.Argument{
					{
						Name:  "arg",
						Value: "test",
					},
				},
			},
		},
		Fields: []schema.Field{
			{
				Name: "field2",
				Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
			},
			{
				Name: "field",
				Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
			},
		},
	}

	assert.True(t, o.Equals(&schema.Object{
		Name: "Object",
		Directives: []schema.Directive{
			{
				Name: "directive",
				Arguments: []schema.Argument{
					{
						Name:  "arg",
						Value: "test",
					},
				},
			},
			{
				Name: "directive",
				Arguments: []schema.Argument{
					{
						Name:  "arg2",
						Value: "test",
					},
				},
			},
		},
		Fields: []schema.Field{
			{
				Name: "field",
				Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
			},
			{
				Name: "field2",
				Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
			},
		},
	}))

	assert.False(t, o.Equals(&schema.Object{
		Name: "Object",
		Directives: []schema.Directive{
			{
				Name: "directive",
				Arguments: []schema.Argument{
					{
						Name:  "arg",
						Value: "test",
					},
				},
			},
			{
				Name: "directive",
				Arguments: []schema.Argument{
					{
						Name:  "arg2",
						Value: "test",
					},
				},
			},
		},
		Fields: []schema.Field{
			{
				Name: "field",
				Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
			},
			{
				Name: "field3",
				Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
			},
		},
	}))
}
