package validator

import (
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/stretchr/testify/mock"
)

type MockFieldValidator struct {
	mock.Mock
}

func (v *MockFieldValidator) Validate(field *schema.Field, object *schema.Object, schema *schema.Schema, customScalarTypes []string) error {
	return v.Called(field, object, schema, customScalarTypes).Error(0)
}
