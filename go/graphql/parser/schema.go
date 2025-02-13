package parser

import (
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/source"
)

func GetSchemaFromString(s string) (*schema.Schema, error) {
	document, err := getDocumentFromSchema(s)
	if err != nil {
		return nil, err
	}

	objects, err := getObjectsFromDocument(document)
	if err != nil {
		return nil, err
	}

	enums, err := getEnumsFromDocument(document)
	if err != nil {
		return nil, err
	}

	scalars, err := getScalarsFromDocument(document)
	if err != nil {
		return nil, err
	}

	return &schema.Schema{
		Objects: objects,
		Enums:   enums,
		Scalars: scalars,
	}, nil
}

func getDocumentFromSchema(schema string) (*ast.Document, error) {
	return parser.Parse(parser.ParseParams{
		Source: source.NewSource(&source.Source{
			Body: []byte(schema),
		}),
		Options: parser.ParseOptions{},
	})
}

func getObjectsFromDocument(document *ast.Document) ([]schema.Object, error) {
	var objects []schema.Object

	for _, definition := range document.Definitions {
		switch definition.GetKind() {
		case kinds.ObjectDefinition:
			object, err := GetObjectFromDefinition(definition.(*ast.ObjectDefinition))
			if err != nil {
				return nil, err
			}

			objects = append(objects, *object)
		}
	}

	return objects, nil
}

func getEnumsFromDocument(document *ast.Document) ([]schema.Enum, error) {
	var enums []schema.Enum

	for _, definition := range document.Definitions {
		switch definition.GetKind() {
		case kinds.EnumDefinition:
			enum, err := GetEnumFromDefinition(definition.(*ast.EnumDefinition))
			if err != nil {
				return nil, err
			}

			enums = append(enums, *enum)
		}
	}

	return enums, nil
}

func getScalarsFromDocument(document *ast.Document) ([]schema.Scalar, error) {
	var scalars []schema.Scalar

	for _, definition := range document.Definitions {
		switch definition.GetKind() {
		case kinds.ScalarDefinition:
			enum, err := GetScalarFromDefinition(definition.(*ast.ScalarDefinition))
			if err != nil {
				return nil, err
			}

			scalars = append(scalars, *enum)
		}
	}

	return scalars, nil
}
