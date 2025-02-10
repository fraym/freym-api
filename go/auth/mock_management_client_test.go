package auth_test

import (
	"testing"

	"github.com/fraym/freym-api/go/auth"
	"github.com/stretchr/testify/assert"
)

func TestMockManagementClient(t *testing.T) {
	assert.Implements(t, new(auth.ManagementClient), new(auth.MockManagementClient))
}
