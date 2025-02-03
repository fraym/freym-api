package streams_test

import (
	"testing"

	"github.com/fraym/freym-api/go/streams"
	"github.com/stretchr/testify/assert"
)

func TestMockClient(t *testing.T) {
	assert.Implements(t, new(streams.Client), new(streams.MockClient))
}
