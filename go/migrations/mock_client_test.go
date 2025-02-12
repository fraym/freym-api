package migrations_test

import (
	"testing"

	"github.com/fraym/freym-api/go/migrations"
	"github.com/stretchr/testify/assert"
)

func TestMockClient(t *testing.T) {
	assert.Implements(t, new(migrations.Client), new(migrations.MockClient))
}
