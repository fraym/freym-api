package schema_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func TestArgument_GetString_Int(t *testing.T) {
	a := schema.Argument{
		Name:  "arg",
		Kind:  kinds.IntValue,
		Value: 123,
	}

	expectedString := "arg: 123"

	schemaString, err := a.GetString()
	assert.Equal(t, expectedString, schemaString)
	assert.NoError(t, err)
}

func TestArgument_GetString_Float(t *testing.T) {
	a := schema.Argument{
		Name:  "arg",
		Kind:  kinds.FloatValue,
		Value: 123.45,
	}

	expectedString := "arg: 123.45"

	schemaString, err := a.GetString()
	assert.Equal(t, expectedString, schemaString)
	assert.NoError(t, err)
}

func TestArgument_GetString_Bool(t *testing.T) {
	a := schema.Argument{
		Name:  "arg",
		Kind:  kinds.BooleanValue,
		Value: true,
	}

	expectedString := "arg: true"

	schemaString, err := a.GetString()
	assert.Equal(t, expectedString, schemaString)
	assert.NoError(t, err)
}

func TestArgument_GetString_List(t *testing.T) {
	a := schema.Argument{
		Name: "arg",
		Kind: kinds.ListValue,
		Value: []any{"string", 123, 123.45, true, map[string]any{
			"d": true,
			"e": "test",
		}},
	}

	expectedString := "arg: [\"string\",123,123.45,true,{ d: true, e: \"test\" }]"

	schemaString, err := a.GetString()
	assert.Equal(t, expectedString, schemaString)
	assert.NoError(t, err)
}

func TestArgument_GetString_Object(t *testing.T) {
	a := schema.Argument{
		Name: "arg",
		Kind: kinds.ObjectValue,
		Value: map[string]any{
			"a": "string",
			"b": 123.45,
			"c": map[string]any{
				"d": true,
			},
		},
	}

	expectedString := "arg: { a: \"string\", b: 123.45, c: { d: true } }"

	schemaString, err := a.GetString()
	assert.Equal(t, expectedString, schemaString)
	assert.NoError(t, err)
}

func TestArgument_Equals(t *testing.T) {
	a := schema.Argument{
		Name: "arg",
		Kind: kinds.ObjectValue,
		Value: map[string]any{
			"a": "string",
			"b": 123.45,
			"c": map[string]any{
				"d": true,
				"a": map[string]any{
					"e2": 123,
					"e1": 12,
				},
			},
		},
	}

	assert.True(t, a.Equals(&schema.Argument{
		Name: "arg",
		Kind: kinds.ObjectValue,
		Value: map[string]any{
			"a": "string",
			"c": map[string]any{
				"d": true,
				"a": map[string]any{
					"e1": 12,
					"e2": 123,
				},
			},
			"b": 123.45,
		},
	}))

	assert.False(t, a.Equals(&schema.Argument{
		Name:  "arg",
		Kind:  kinds.ObjectValue,
		Value: "stringValue",
	}))
}
