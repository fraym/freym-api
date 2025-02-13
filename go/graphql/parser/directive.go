package parser

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/ast"
)

func GetDirectivesFromObjectDefinition(definition *ast.ObjectDefinition) ([]schema.Directive, error) {
	var directives []schema.Directive

	for _, directive := range definition.Directives {
		arguments, err := GetArgumentsOfDirective(directive)
		if err != nil {
			return nil, err
		}

		if directive == nil || directive.Name == nil || len(directive.Name.Value) == 0 {
			return nil, fmt.Errorf("cannot parse directive without name on object `%s`", definition.Name.Value)
		}

		directives = append(directives, schema.Directive{
			Name:      directive.Name.Value,
			Arguments: arguments,
		})
	}

	return directives, nil
}

func GetDirectivesFromFieldDefinition(field *ast.FieldDefinition) ([]schema.Directive, error) {
	var directives []schema.Directive

	for _, directive := range field.Directives {
		arguments, err := GetArgumentsOfDirective(directive)
		if err != nil {
			return nil, err
		}

		if directive == nil || directive.Name == nil || len(directive.Name.Value) == 0 {
			return nil, fmt.Errorf("cannot parse directive without name on field `%s`", field.Name.Value)
		}

		directives = append(directives, schema.Directive{
			Name:      directive.Name.Value,
			Arguments: arguments,
		})
	}

	return directives, nil
}

func GetDirectivesFromScalarDefinition(scalar *ast.ScalarDefinition) ([]schema.Directive, error) {
	var directives []schema.Directive

	for _, directive := range scalar.Directives {
		arguments, err := GetArgumentsOfDirective(directive)
		if err != nil {
			return nil, err
		}

		if directive == nil || directive.Name == nil || len(directive.Name.Value) == 0 {
			return nil, fmt.Errorf("cannot parse directive without name on scalar `%s`", scalar.Name.Value)
		}

		directives = append(directives, schema.Directive{
			Name:      directive.Name.Value,
			Arguments: arguments,
		})
	}

	return directives, nil
}
