package validator

import (
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/stretchr/testify/mock"
)

type MockDirectiveValidator struct {
	mock.Mock
}

func (v *MockDirectiveValidator) ValidateObjectDirective(directive *schema.Directive) error {
	return v.Called(directive).Error(0)
}

func (v *MockDirectiveValidator) ValidateFieldDirective(directive *schema.Directive, objectDirectiveNames []string, fieldType types.Comparable) error {
	return v.Called(directive, objectDirectiveNames, fieldType).Error(0)
}
