package graphql

import (
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/validator"
)

type SchemaValidator struct {
	object    *validator.ObjectValidator
	directive *validator.DirectiveValidator
	field     *validator.FieldValidator
}

func NewSchemaValidator(
	objectConfig *validator.ObjectConfig,
	directiveConfig *validator.DirectiveConfig,
	fieldConfig *validator.FieldConfig,
) *SchemaValidator {
	directive := validator.NewDirectiveValidator(directiveConfig)
	field := validator.NewFieldValidator(fieldConfig, directive)
	object := validator.NewObjectValidator(objectConfig, directive, field)

	return &SchemaValidator{
		object:    object,
		directive: directive,
		field:     field,
	}
}

func (v *SchemaValidator) Validate(schema *schema.Schema, customScalarTypes []string) error {
	for _, object := range schema.Objects {
		if err := v.object.Validate(&object, schema, customScalarTypes); err != nil {
			return err
		}
	}

	return nil
}
