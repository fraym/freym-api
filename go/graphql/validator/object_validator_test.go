package validator_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/validator"
	"github.com/stretchr/testify/assert"
)

func TestObjectValidator_Validate_MissingRequiredDirective(t *testing.T) {
	directiveValidator := &validator.MockDirectiveValidator{}
	fieldValidator := &validator.MockFieldValidator{}
	v := validator.NewObjectValidator(&validator.ObjectConfig{
		RequireObjectDirective: true,
	}, directiveValidator, fieldValidator)

	expectedCustomScalarTypes := []string{}
	expectedSchema := &schema.Schema{}

	err := v.Validate(&schema.Object{
		Name: "object",
	}, expectedSchema, expectedCustomScalarTypes)
	assert.EqualError(t, err, "object `object` is required to use at least one object directive")

	directiveValidator.AssertExpectations(t)
	fieldValidator.AssertExpectations(t)
}

func TestObjectValidator_Validate_MissingFields(t *testing.T) {
	directiveValidator := &validator.MockDirectiveValidator{}
	fieldValidator := &validator.MockFieldValidator{}
	v := validator.NewObjectValidator(&validator.ObjectConfig{}, directiveValidator, fieldValidator)

	expectedCustomScalarTypes := []string{}
	expectedSchema := &schema.Schema{}

	err := v.Validate(&schema.Object{
		Name: "object",
	}, expectedSchema, expectedCustomScalarTypes)
	assert.EqualError(t, err, "object `object` does not contain any fields")

	directiveValidator.AssertExpectations(t)
	fieldValidator.AssertExpectations(t)
}

func TestObjectValidator_Validate_FieldsAndDirective(t *testing.T) {
	directiveValidator := &validator.MockDirectiveValidator{}
	fieldValidator := &validator.MockFieldValidator{}
	v := validator.NewObjectValidator(&validator.ObjectConfig{}, directiveValidator, fieldValidator)

	expecrtedDirective := schema.Directive{Name: "directive"}
	expectedField := schema.Field{Name: "field"}
	expectedCustomScalarTypes := []string{}
	expectedSchema := &schema.Schema{}
	expectedObject := &schema.Object{
		Name:       "object",
		Directives: []schema.Directive{expecrtedDirective},
		Fields:     []schema.Field{expectedField},
	}

	directiveValidator.On("ValidateObjectDirective", &expecrtedDirective).Return(nil)
	fieldValidator.On("Validate", &expectedField, expectedObject, expectedSchema, expectedCustomScalarTypes).Return(nil)

	err := v.Validate(expectedObject, expectedSchema, expectedCustomScalarTypes)
	assert.NoError(t, err)

	directiveValidator.AssertExpectations(t)
	fieldValidator.AssertExpectations(t)
}
