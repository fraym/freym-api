package parser

import (
	"fmt"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/ast"
)

func GetArgumentsOfDirective(directive *ast.Directive) ([]schema.Argument, error) {
	var arguments []schema.Argument

	for _, argument := range directive.Arguments {
		value, err := GetValue(argument.Value)
		if err != nil {
			return nil, err
		}

		if argument == nil || argument.Name == nil || len(argument.Name.Value) == 0 {
			return nil, fmt.Errorf("cannot parse argument without name on directive `%s`", directive.Name.Value)
		}

		arguments = append(arguments, schema.Argument{
			Name:  argument.Name.Value,
			Value: value,
			Kind:  argument.Value.GetKind(),
		})
	}

	return arguments, nil
}
