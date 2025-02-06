package projections_test

import (
	"testing"

	"github.com/fraym/freym-api/go/projections"
	"github.com/stretchr/testify/assert"
)

func TestMockManagementClient(t *testing.T) {
	assert.Implements(t, new(projections.ManagementClient), new(projections.MockManagementClient))
}
