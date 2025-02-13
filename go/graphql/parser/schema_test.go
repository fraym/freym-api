package parser_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func TestGetSchemaFromString(t *testing.T) {
	schemaString := `
		scalar ScalarType

		enum EnumType {
			V1
			V2
			V3
		}

		type Object {
			field: String!
			scalar: ScalarName!
			ref: Reference!
			enum: EnumType!
		}

		type Reference {
			field: String
		}
	`

	s, err := parser.GetSchemaFromString(schemaString)

	assert.NoError(t, err)
	assert.Equal(t, &schema.Schema{
		Objects: []schema.Object{
			{
				Name: "Object",
				Fields: []schema.Field{
					{
						Name: "field",
						Type: types.Comparable{kinds.NonNull, kinds.Named, "String"},
					},
					{
						Name: "scalar",
						Type: types.Comparable{kinds.NonNull, kinds.Named, "ScalarName"},
					},
					{
						Name: "ref",
						Type: types.Comparable{kinds.NonNull, kinds.Named, "Reference"},
					},
					{
						Name: "enum",
						Type: types.Comparable{kinds.NonNull, kinds.Named, "EnumType"},
					},
				},
			},
			{
				Name: "Reference",
				Fields: []schema.Field{
					{
						Name: "field",
						Type: types.Comparable{kinds.Named, "String"},
					},
				},
			},
		},
		Enums: []schema.Enum{
			{
				Name:   "EnumType",
				Values: []string{"V1", "V2", "V3"},
			},
		},
		Scalars: []schema.Scalar{
			{
				Name: "ScalarType",
			},
		},
	}, s)
}
