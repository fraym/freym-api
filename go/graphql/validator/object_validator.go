package validator

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
)

type ObjectConfig struct {
	// set to true if at least one directive is required
	RequireObjectDirective bool
}

type ObjectValidatorDirectiveValidator interface {
	ValidateObjectDirective(directive *schema.Directive) error
	ValidateFieldDirective(directive *schema.Directive, objectDirectiveNames []string, fieldType types.Comparable) error
}

type ObjectValidatorFieldValidator interface {
	Validate(field *schema.Field, object *schema.Object, schema *schema.Schema, customScalarTypes []string) error
}

type ObjectValidator struct {
	config             *ObjectConfig
	directiveValidator ObjectValidatorDirectiveValidator
	fieldValidator     ObjectValidatorFieldValidator
}

func NewObjectValidator(
	config *ObjectConfig,
	directiveValidator ObjectValidatorDirectiveValidator,
	fieldValidator ObjectValidatorFieldValidator,
) *ObjectValidator {
	if config == nil {
		config = &ObjectConfig{}
	}

	return &ObjectValidator{
		config:             config,
		directiveValidator: directiveValidator,
		fieldValidator:     fieldValidator,
	}
}

func (v *ObjectValidator) Validate(object *schema.Object, schema *schema.Schema, customScalarTypes []string) error {
	if err := v.validateDirectives(object); err != nil {
		return err
	}

	return v.validateFields(object, schema, customScalarTypes)
}

func (v *ObjectValidator) validateDirectives(object *schema.Object) error {
	if v.config.RequireObjectDirective && len(object.Directives) == 0 {
		return fmt.Errorf("object `%s` is required to use at least one object directive", object.Name)
	}

	for _, directive := range object.Directives {
		if err := v.directiveValidator.ValidateObjectDirective(&directive); err != nil {
			return err
		}
	}

	return nil
}

func (v *ObjectValidator) validateFields(
	object *schema.Object,
	schema *schema.Schema,
	customScalarTypes []string,
) error {
	if len(object.Fields) == 0 {
		return fmt.Errorf("object `%s` does not contain any fields", object.Name)
	}

	for _, field := range object.Fields {
		if err := v.fieldValidator.Validate(&field, object, schema, customScalarTypes); err != nil {
			return err
		}
	}

	return nil
}
