package deployments_test

import (
	"testing"

	"github.com/fraym/freym-api/go/deployments"
	"github.com/stretchr/testify/assert"
)

func TestMockClient(t *testing.T) {
	assert.Implements(t, new(deployments.Client), new(deployments.MockClient))
}
