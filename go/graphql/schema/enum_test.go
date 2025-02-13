package schema_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/stretchr/testify/assert"
)

func TestEnum_GetString(t *testing.T) {
	e := schema.Enum{
		Name:   "Enum",
		Values: []string{"V1", "V2"},
	}

	expectedString := "enum Enum {\n\tV1\n\tV2\n}"

	schemaString, err := e.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, schemaString)
}
