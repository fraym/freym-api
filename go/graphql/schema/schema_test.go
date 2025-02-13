package schema_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/stretchr/testify/assert"
)

func TestSchema_GetString(t *testing.T) {
	s := schema.Schema{
		Objects: []schema.Object{
			{
				Name: "Object1",
			},
			{
				Name: "Object2",
			},
		},
		Enums: []schema.Enum{
			{
				Name: "Enum",
			},
		},
		Scalars: []schema.Scalar{
			{
				Name: "Scalar",
			},
		},
	}

	expectedString := "enum Enum {\n}\ntype Object1 {\n}\ntype Object2 {\n}\nscalar Scalar\n"

	schemaString, err := s.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, schemaString)
}
