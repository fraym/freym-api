package validator_test

import (
	"errors"
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/fraym/freym-api/go/graphql/validator"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type typeValidator struct {
	mock.Mock
}

func (v *typeValidator) validate(field *schema.Field, object *schema.Object, schema *schema.Schema, customScalarTypes []string) error {
	return v.Called(field, object, schema, customScalarTypes).Error(0)
}

func prepareFieldValidator() (*validator.FieldValidator, *validator.MockDirectiveValidator, *typeValidator) {
	typeValidator := &typeValidator{}
	directiveValidator := &validator.MockDirectiveValidator{}
	return validator.NewFieldValidator(&validator.FieldConfig{
		AllowedTypes: []types.Comparable{
			{kinds.Named, "String"},
		},
		TypeValidator: typeValidator.validate,
	}, directiveValidator), directiveValidator, typeValidator
}

func TestFieldValidator_Validate_InvalidType(t *testing.T) {
	v, directiveValidator, typeValidator := prepareFieldValidator()

	expectedComparable := types.Comparable{kinds.Named, "invalid"}
	expectedErr := errors.New("test")
	expectedField := &schema.Field{
		Name: "field",
		Type: expectedComparable,
	}
	expectedObject := &schema.Object{}
	expectedSchema := &schema.Schema{}
	expectedCustomScalarTypes := []string{}

	typeValidator.On("validate", expectedField, expectedObject, expectedSchema, expectedCustomScalarTypes).Return(expectedErr).Once()

	err := v.Validate(expectedField, expectedObject, expectedSchema, expectedCustomScalarTypes)
	assert.ErrorIs(t, err, expectedErr)

	typeValidator.AssertExpectations(t)
	directiveValidator.AssertExpectations(t)
}

func TestFieldValidator_Validate_CorrectType(t *testing.T) {
	v, directiveValidator, typeValidator := prepareFieldValidator()

	expectedComparable := types.Comparable{kinds.Named, "String"}
	expectedObject := &schema.Object{}
	expectedSchema := &schema.Schema{}
	expectedCustomScalarTypes := []string{}

	field := &schema.Field{
		Name: "field",
		Type: expectedComparable,
	}
	err := v.Validate(field, expectedObject, expectedSchema, expectedCustomScalarTypes)
	assert.NoError(t, err)

	typeValidator.AssertExpectations(t)
	directiveValidator.AssertExpectations(t)
}

func TestFieldValidator_Validate_InvalidDirective(t *testing.T) {
	v, directiveValidator, typeValidator := prepareFieldValidator()

	expectedDirective := &schema.Directive{
		Name: "directive",
	}
	expectedComparable := types.Comparable{kinds.Named, "String"}
	expectedObjectDirectives := []string(nil)
	expectedObject := &schema.Object{}
	expectedSchema := &schema.Schema{}
	expectedCustomScalarTypes := []string{}
	expectedErr := errors.New("test")

	directiveValidator.On("ValidateFieldDirective", expectedDirective, expectedObjectDirectives, expectedComparable).Return(expectedErr)

	field := &schema.Field{
		Name: "field",
		Type: expectedComparable,
		Directives: []schema.Directive{
			*expectedDirective,
		},
	}
	err := v.Validate(field, expectedObject, expectedSchema, expectedCustomScalarTypes)
	assert.ErrorIs(t, err, expectedErr)

	typeValidator.AssertExpectations(t)
	directiveValidator.AssertExpectations(t)
}
