package crud_test

import (
	"testing"

	"github.com/fraym/freym-api/go/crud"
	"github.com/stretchr/testify/assert"
)

func TestMockDeliveryClient(t *testing.T) {
	assert.Implements(t, new(crud.DeliveryClient), new(crud.MockDeliveryClient))
}
