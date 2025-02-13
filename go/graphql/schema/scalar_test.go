package schema_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/stretchr/testify/assert"
)

func TestScalar_GetString(t *testing.T) {
	s := schema.Scalar{
		Name: "Scalar",
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

	expectedString := "scalar Scalar @directive @directive2"

	schemaString, err := s.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, schemaString)
}

func TestScalar_Equals(t *testing.T) {
	s := schema.Scalar{
		Name: "Scalar",
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

	assert.True(t, s.Equals(&schema.Scalar{
		Name: "Scalar",
		Directives: []schema.Directive{
			{
				Name:      "directive2",
				Arguments: []schema.Argument{},
			},
			{
				Name:      "directive",
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
		},
	}))

	assert.False(t, s.Equals(&schema.Scalar{
		Name: "Scalar",
		Directives: []schema.Directive{
			{
				Name:      "directive2",
				Arguments: []schema.Argument{},
			},
			{
				Name:      "directive",
				Arguments: []schema.Argument{},
			},
			{
				Name: "directive2",
				Arguments: []schema.Argument{
					{
						Name:  "a1",
						Value: "test",
					},
				},
			},
		},
	}))
}
