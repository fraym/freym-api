package parser_test

import (
	"strconv"
	"testing"

	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/stretchr/testify/assert"
)

func TestGetValue_IntValue(t *testing.T) {
	expectedValue := 123

	valueData := ast.NewIntValue(&ast.IntValue{
		Value: strconv.Itoa(expectedValue),
	})

	value, err := parser.GetValue(valueData)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)
}

func TestGetValue_FloatValue(t *testing.T) {
	expectedValue := 12.34

	valueData := ast.NewFloatValue(&ast.FloatValue{
		Value: strconv.FormatFloat(expectedValue, 'f', 2, 64),
	})

	value, err := parser.GetValue(valueData)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)
}

func TestGetValue_StringValue(t *testing.T) {
	expectedValue := "test"

	valueData := ast.NewStringValue(&ast.StringValue{
		Value: expectedValue,
	})

	value, err := parser.GetValue(valueData)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)
}

func TestGetValue_EnumValue(t *testing.T) {
	expectedValue := "test"

	valueData := ast.NewEnumValue(&ast.EnumValue{
		Value: expectedValue,
	})

	value, err := parser.GetValue(valueData)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)
}

func TestGetValue_BoolValue(t *testing.T) {
	expectedValue := true

	valueData := ast.NewBooleanValue(&ast.BooleanValue{
		Value: expectedValue,
	})

	value, err := parser.GetValue(valueData)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)
}

func TestGetValue_ListValue(t *testing.T) {
	expectedStringValue := "test"
	expectedIntValue := 123

	valueData := ast.NewListValue(&ast.ListValue{
		Values: []ast.Value{
			ast.NewStringValue(&ast.StringValue{
				Value: expectedStringValue,
			}),
			ast.NewIntValue(&ast.IntValue{
				Value: strconv.Itoa(expectedIntValue),
			}),
		},
	})

	value, err := parser.GetValue(valueData)
	values := value.([]any)
	assert.NoError(t, err)
	assert.Equal(t, expectedStringValue, values[0])
	assert.Equal(t, expectedIntValue, values[1])
}

func TestGetValue_ObjectValue(t *testing.T) {
	expectedStringValue := "test"
	expectedIntValue := 123

	valueData := ast.NewObjectValue(&ast.ObjectValue{
		Fields: []*ast.ObjectField{
			ast.NewObjectField(&ast.ObjectField{
				Name: ast.NewName(&ast.Name{
					Value: "stringKey",
				}),
				Value: ast.NewStringValue(&ast.StringValue{
					Value: expectedStringValue,
				}),
			}),
			ast.NewObjectField(&ast.ObjectField{
				Name: ast.NewName(&ast.Name{
					Value: "intKey",
				}),
				Value: ast.NewIntValue(&ast.IntValue{
					Value: strconv.Itoa(expectedIntValue),
				}),
			}),
		},
	})

	value, err := parser.GetValue(valueData)
	values := value.(map[string]any)
	assert.NoError(t, err)
	assert.Len(t, values, 2)
	assert.Equal(t, expectedStringValue, values["stringKey"])
	assert.Equal(t, expectedIntValue, values["intKey"])
}
