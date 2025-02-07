package crud_test

import (
	"testing"

	"github.com/fraym/freym-api/go/crud"
	"github.com/stretchr/testify/assert"
)

func TestMockManagementClient(t *testing.T) {
	assert.Implements(t, new(crud.ManagementClient), new(crud.MockManagementClient))
}
