package projections_test

import (
	"testing"

	"github.com/fraym/freym-api/go/projections"
	"github.com/stretchr/testify/assert"
)

func TestMockDeliveryClient(t *testing.T) {
	assert.Implements(t, new(projections.DeliveryClient), new(projections.MockDeliveryClient))
}
