package validator

import (
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/stretchr/testify/mock"
)

type MockObjectValidator struct {
	mock.Mock
}

func (v *MockObjectValidator) Validate(obj *schema.Object, customScalarTypes []string) error {
	return v.Called(obj, customScalarTypes).Error(0)
}
