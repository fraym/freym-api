package sync_test

import (
	"testing"

	"github.com/fraym/freym-api/go/sync"
	"github.com/stretchr/testify/assert"
)

type ServiceClient struct{}

func TestMockClient(t *testing.T) {
	assert.Implements(t, new(sync.Client[*ServiceClient]), new(sync.MockClient[*ServiceClient]))
}
