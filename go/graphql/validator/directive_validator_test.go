package validator_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/fraym/freym-api/go/graphql/validator"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func prepareObjectDirectiveValidator() *validator.DirectiveValidator {
	return validator.NewDirectiveValidator(&validator.DirectiveConfig{
		AllowedObjectDirectives: map[string]validator.ObjectDirectiveValidationOptions{
			"directive": {
				RequiredArgumentKinds: map[string]string{
					"required": kinds.StringValue,
				},
				OptionalArgumentKinds: map[string]string{
					"optional": kinds.IntValue,
					"all":      "*",
				},
			},
		},
	})
}

func prepareFieldDirectiveValidator() *validator.DirectiveValidator {
	return validator.NewDirectiveValidator(&validator.DirectiveConfig{
		AllowedFieldDirectives: map[string]validator.FieldDirectiveOptions{
			"directive": {
				AllowedFieldType: []types.Comparable{
					{"String", kinds.Named},
				},
				RequiredArgumentKinds: map[string]string{
					"required": kinds.StringValue,
				},
				OptionalArgumentKinds: map[string]string{
					"optional": kinds.IntValue,
				},
			},
			"oneDirective": {
				RequireOneObjectDirectiveOf: []string{"directive1", "directive2"},
				AllowedFieldType: []types.Comparable{
					{"String", kinds.Named},
				},
			},
			"allDirectives": {
				RequireAllObjectDirectivesOf: []string{"directive1", "directive2"},
				AllowedFieldType: []types.Comparable{
					{"String", kinds.Named},
				},
			},
			"allFields": {
				AllowAllFieldTypes: true,
			},
			"allKinds": {
				AllowAllFieldTypes: true,
				RequiredArgumentKinds: map[string]string{
					"required": "*",
				},
			},
		},
	})
}

func TestValidateObjectDirective_MissingRequiredArgument(t *testing.T) {
	v := prepareObjectDirectiveValidator()
	directive := &schema.Directive{
		Name: "directive",
	}

	err := v.ValidateObjectDirective(directive)
	assert.Error(t, err)
}

func TestValidateObjectDirective_WrongArgumentForRequiredArgument(t *testing.T) {
	v := prepareObjectDirectiveValidator()
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Value: 123,
				Kind:  kinds.IntValue,
			},
		},
	}

	err := v.ValidateObjectDirective(directive)
	assert.Error(t, err)
}

func TestValidateObjectDirective_NoOptionalArguments(t *testing.T) {
	v := prepareObjectDirectiveValidator()
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Value: "string",
				Kind:  kinds.StringValue,
			},
		},
	}

	err := v.ValidateObjectDirective(directive)
	assert.NoError(t, err)
}

func TestValidateObjectDirective_WrongArgumentForOptionalArgument(t *testing.T) {
	v := prepareObjectDirectiveValidator()
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Value: "string",
				Kind:  kinds.StringValue,
			},
			{
				Name:  "optional",
				Value: "string",
				Kind:  kinds.StringValue,
			},
		},
	}

	err := v.ValidateObjectDirective(directive)
	assert.Error(t, err)
}

func TestValidateObjectDirective_CorrectDirective(t *testing.T) {
	v := prepareObjectDirectiveValidator()
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Value: "string",
				Kind:  kinds.StringValue,
			},
			{
				Name:  "optional",
				Value: 123,
				Kind:  kinds.IntValue,
			},
		},
	}

	err := v.ValidateObjectDirective(directive)
	assert.NoError(t, err)
}

func TestValidateObjectDirective_InvalidArgument(t *testing.T) {
	v := prepareObjectDirectiveValidator()
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Value: "string",
				Kind:  kinds.StringValue,
			},
			{
				Name:  "optional",
				Value: 123,
				Kind:  kinds.IntValue,
			},
			{
				Name:  "invalid",
				Value: 123,
				Kind:  kinds.IntValue,
			},
		},
	}

	err := v.ValidateObjectDirective(directive)
	assert.Error(t, err)
}

func TestValidateObjectDirective_AllKinds(t *testing.T) {
	v := prepareObjectDirectiveValidator()
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Value: "string",
				Kind:  kinds.StringValue,
			},
			{
				Name:  "all",
				Value: "string",
				Kind:  kinds.StringValue,
			},
		},
	}

	err := v.ValidateObjectDirective(directive)
	assert.NoError(t, err)

	directive.Arguments[1].Kind = kinds.IntValue
	directive.Arguments[1].Value = 123

	err = v.ValidateObjectDirective(directive)
	assert.NoError(t, err)
}

func TestValidateFieldDirective_InvalidFieldType(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{}
	objectDirectives := []string{}
	directive := &schema.Directive{
		Name: "allFields",
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.NoError(t, err)
}

func TestValidateFieldDirective_AllFieldTypes(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "directive",
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.Error(t, err)
}

func TestValidateFieldDirective_MissingArgument(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "directive",
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.Error(t, err)
}

func TestValidateFieldDirective_WrongArgumentForRequiredArgument(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Kind:  kinds.IntValue,
				Value: 123,
			},
		},
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.Error(t, err)
}

func TestValidateFieldDirective_NoOptionalArguments(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Kind:  kinds.StringValue,
				Value: "string",
			},
		},
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.NoError(t, err)
}

func TestValidateFieldDirective_AllArgumentKinds(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{}
	directive := &schema.Directive{
		Name: "allKinds",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Kind:  kinds.StringValue,
				Value: "string",
			},
		},
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.NoError(t, err)

	directive.Arguments[0].Kind = kinds.IntValue
	directive.Arguments[0].Value = 123

	err = v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.NoError(t, err)
}

func TestValidateFieldDirective_WrongArgumentForOptionalArgument(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Kind:  kinds.StringValue,
				Value: "string",
			},
			{
				Name:  "optional",
				Kind:  kinds.StringValue,
				Value: "string",
			},
		},
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.Error(t, err)
}

func TestValidateFieldDirective_CorrectDirective(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Kind:  kinds.StringValue,
				Value: "string",
			},
			{
				Name:  "optional",
				Kind:  kinds.IntValue,
				Value: 123,
			},
		},
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.NoError(t, err)
}

func TestValidateFieldDirective_InvalidArgument(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "directive",
		Arguments: []schema.Argument{
			{
				Name:  "required",
				Kind:  kinds.StringValue,
				Value: "string",
			},
			{
				Name:  "optional",
				Kind:  kinds.IntValue,
				Value: 123,
			},
			{
				Name:  "invalid",
				Kind:  kinds.IntValue,
				Value: 123,
			},
		},
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.Error(t, err)
}

func TestValidateFieldDirective_MissingObjectDirective(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{}
	directive := &schema.Directive{
		Name: "oneDirective",
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.Error(t, err)
}

func TestValidateFieldDirective_NoMissingObjectDirective(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "oneDirective",
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.NoError(t, err)
}

func TestValidateFieldDirective_RequireAllMissingObjectDirective(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1"}
	directive := &schema.Directive{
		Name: "allDirectives",
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.Error(t, err)
}

func TestValidateFieldDirective_RequireAllNoMissingObjectDirective(t *testing.T) {
	v := prepareFieldDirectiveValidator()

	fieldType := types.Comparable{"String", kinds.Named}
	objectDirectives := []string{"directive1", "directive2"}
	directive := &schema.Directive{
		Name: "allDirectives",
	}

	err := v.ValidateFieldDirective(directive, objectDirectives, fieldType)
	assert.NoError(t, err)
}
