package validator

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/samber/lo"
)

type ObjectDirectiveValidationOptions struct {
	// list of required arguments and their kinds (use * to allow all kinds)
	RequiredArgumentKinds map[string]string
	// list of optional arguments and their kinds (use * to allow all kinds)
	OptionalArgumentKinds map[string]string
}

type FieldDirectiveOptions struct {
	// list of field types that are allowed for this directive
	AllowedFieldType []types.Comparable
	// allow this directive on all field types
	AllowAllFieldTypes bool
	// list of required arguments and their kinds (use * to allow all kinds)
	RequiredArgumentKinds map[string]string
	// list of optional arguments and their kinds (use * to allow all kinds)
	OptionalArgumentKinds map[string]string
	// one of these directives in this list is required to be defined on the object
	RequireOneObjectDirectiveOf []string
	// all of these directives in this list are required to be defined on the object
	RequireAllObjectDirectivesOf []string
}

type DirectiveConfig struct {
	// directives that are allowed to be defined on objects
	AllowedObjectDirectives map[string]ObjectDirectiveValidationOptions
	// directives that are allowed to be defined on fields
	AllowedFieldDirectives map[string]FieldDirectiveOptions
}

type DirectiveValidator struct {
	config *DirectiveConfig
}

func NewDirectiveValidator(
	config *DirectiveConfig,
) *DirectiveValidator {
	return &DirectiveValidator{
		config: config,
	}
}

func (v *DirectiveValidator) ValidateObjectDirective(directive *schema.Directive) error {
	options, ok := v.config.AllowedObjectDirectives[directive.Name]
	if !ok {
		return fmt.Errorf("object directive `%s` is not allowed", directive.Name)
	}

	return validateDirectiveArguments(directive, options.RequiredArgumentKinds, options.OptionalArgumentKinds)
}

func (v *DirectiveValidator) ValidateFieldDirective(
	directive *schema.Directive,
	objectDirectives []string,
	fieldType types.Comparable,
) error {
	name := directive.Name
	options, ok := v.config.AllowedFieldDirectives[directive.Name]
	if !ok {
		return fmt.Errorf("field directive `%s` is not allowed", name)
	}

	validFieldType := options.AllowAllFieldTypes

	if !options.AllowAllFieldTypes {
		for _, allowed := range options.AllowedFieldType {
			if allowed.Equals(fieldType) {
				validFieldType = true
				break
			}
		}
	}

	if !validFieldType {
		return fmt.Errorf(
			"invalid field type `%s` for field with directive `%s`, expected one of: %s",
			fieldType,
			name,
			options.AllowedFieldType,
		)
	}

	if len(options.RequireOneObjectDirectiveOf) > 0 && !lo.Some(objectDirectives, options.RequireOneObjectDirectiveOf) {
		return fmt.Errorf(
			"cannot use field directive `%s` without one object directive of: %s",
			name,
			options.RequireOneObjectDirectiveOf,
		)
	}

	if !lo.Every(objectDirectives, options.RequireAllObjectDirectivesOf) {
		return fmt.Errorf(
			"cannot use field directive `%s` without all object directives of: %s",
			name,
			options.RequireAllObjectDirectivesOf,
		)
	}

	return validateDirectiveArguments(directive, options.RequiredArgumentKinds, options.OptionalArgumentKinds)
}

func validateDirectiveArguments(
	directive *schema.Directive,
	requiredArgumentKinds map[string]string,
	optionalArgumentKinds map[string]string,
) error {
	if len(directive.Arguments) < len(requiredArgumentKinds) {
		requiredArgs := lo.Keys(requiredArgumentKinds)
		return fmt.Errorf(
			"directive `%s` does not have all required arguments specified: %s",
			directive.Name,
			requiredArgs,
		)
	}

requirementsLoop:
	for requiredArg, requiredKind := range requiredArgumentKinds {
		for _, arg := range directive.Arguments {
			if arg.Name == requiredArg && (arg.Kind == requiredKind || requiredKind == "*") {
				continue requirementsLoop
			}
		}

		return fmt.Errorf("required argument `%s` of kind `%s` not found on directive `%s`", requiredArg, requiredKind, directive.Name)
	}

	for _, arg := range directive.Arguments {
		name := arg.Name
		if _, ok := requiredArgumentKinds[name]; ok {
			continue
		}

		requiredKind, ok := optionalArgumentKinds[name]
		if !ok {
			allowedKinds := lo.Values(optionalArgumentKinds)
			return fmt.Errorf(
				"invalid argument `%s` on directive `%s`, expected one of: %s",
				name,
				directive.Name,
				allowedKinds,
			)
		}

		if arg.Kind != requiredKind && requiredKind != "*" {
			return fmt.Errorf(
				"invalid kind `%s` for argument `%s` on directive `%s`, `%s` expected",
				arg.Kind,
				name,
				directive.Name,
				requiredKind,
			)
		}
	}

	return nil
}
