package schema_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/stretchr/testify/assert"
)

func TestDirective_GetString(t *testing.T) {
	d := schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "arg",
				Value: "test",
			},
			{
				Name:  "arg2",
				Value: 123,
			},
		},
	}

	expectedString := "@directive(arg: \"test\",arg2: 123)"

	schemaString, err := d.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, schemaString)
}

func TestDirective_GetString_NoArgs(t *testing.T) {
	d := schema.Directive{
		Name:      "directive",
		Arguments: []schema.Argument{},
	}

	expectedString := "@directive"

	schemaString, err := d.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, schemaString)
}
