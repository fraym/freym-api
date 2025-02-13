package validator

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
)

type TypeValidationFunc func(field *schema.Field, object *schema.Object, schema *schema.Schema, customScalarTypes []string) error

type FieldConfig struct {
	// list of allowed types representad as comparable structure
	AllowedTypes []types.Comparable
	// validation func that you can define to validate more complex scenarios like objects
	TypeValidator TypeValidationFunc
}

type FieldValidatorDirectiveValidator interface {
	ValidateFieldDirective(directive *schema.Directive, objectDirectiveNames []string, fieldType types.Comparable) error
}

type FieldValidator struct {
	config             *FieldConfig
	directiveValidator FieldValidatorDirectiveValidator
}

func NewFieldValidator(
	config *FieldConfig,
	directiveValidator FieldValidatorDirectiveValidator,
) *FieldValidator {
	return &FieldValidator{
		config:             config,
		directiveValidator: directiveValidator,
	}
}

func (v *FieldValidator) Validate(field *schema.Field, object *schema.Object, schema *schema.Schema, customScalarTypes []string) error {
	if err := v.validateType(field, object, schema, customScalarTypes); err != nil {
		return err
	}

	var objectDirectiveNames []string

	for _, directive := range object.Directives {
		objectDirectiveNames = append(objectDirectiveNames, directive.Name)
	}

	for _, directive := range field.Directives {
		if err := v.directiveValidator.ValidateFieldDirective(&directive, objectDirectiveNames, field.Type); err != nil {
			return err
		}
	}

	return nil
}

func (v *FieldValidator) validateType(field *schema.Field, object *schema.Object, schema *schema.Schema, customScalarTypes []string) error {
	isValidType := false

	for _, allowedType := range v.config.AllowedTypes {
		if allowedType.Equals(field.Type) {
			isValidType = true
			break
		}
	}

	if !isValidType && v.config.TypeValidator != nil {
		return v.config.TypeValidator(field, object, schema, customScalarTypes)
	}

	if !isValidType {
		return fmt.Errorf("unexpected field type `%s`", field.Type)
	}
	return nil
}
