package migrations

import (
	"encoding/json"

	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/proto/migrations/managementpb"
)

func enumsAndPermissionsToPb(enums []schema.Enum) ([]*managementpb.EnumType, []string, error) {
	var (
		permissions []string
		pbEnums     []*managementpb.EnumType
	)

	for _, enum := range enums {
		if enum.Name == "Permission" {
			permissions = enumToPb(&enum).GetValues()
		} else {
			pbEnums = append(pbEnums, enumToPb(&enum))
		}
	}

	if len(permissions) > 0 {
		return pbEnums, permissions, nil
	}
	return pbEnums, nil, nil
}

func enumToPb(enum *schema.Enum) *managementpb.EnumType {
	return managementpb.EnumType_builder{
		Name:   enum.Name,
		Values: enum.Values,
	}.Build()
}

func objectsToPb(objects []schema.Object) ([]*managementpb.ObjectType, error) {
	var pbObjects []*managementpb.ObjectType

	for _, object := range objects {
		pbObject, err := objectToPb(&object)
		if err != nil {
			return nil, err
		}

		pbObjects = append(pbObjects, pbObject)
	}

	return pbObjects, nil
}

func objectToPb(object *schema.Object) (*managementpb.ObjectType, error) {
	directives, err := directivesToPb(object.Directives)
	if err != nil {
		return nil, err
	}

	fields, err := fieldsToPb(object.Fields)
	if err != nil {
		return nil, err
	}

	return managementpb.ObjectType_builder{
		Name:       object.Name,
		Directives: directives,
		Fields:     fields,
	}.Build(), nil
}

func fieldsToPb(fields []schema.Field) ([]*managementpb.TypeField, error) {
	var pbFields []*managementpb.TypeField

	for _, field := range fields {
		pbField, err := fieldToPb(&field)
		if err != nil {
			return nil, err
		}

		pbFields = append(pbFields, pbField)
	}

	return pbFields, nil
}

func fieldToPb(field *schema.Field) (*managementpb.TypeField, error) {
	directives, err := directivesToPb(field.Directives)
	if err != nil {
		return nil, err
	}

	return managementpb.TypeField_builder{
		Name:       field.Name,
		Type:       field.Type,
		Directives: directives,
	}.Build(), nil
}

func directivesToPb(directives []schema.Directive) ([]*managementpb.TypeDirective, error) {
	var pbDirectives []*managementpb.TypeDirective

	for _, directive := range directives {
		pbDirective, err := directiveToPb(&directive)
		if err != nil {
			return nil, err
		}

		pbDirectives = append(pbDirectives, pbDirective)
	}

	return pbDirectives, nil
}

func directiveToPb(directive *schema.Directive) (*managementpb.TypeDirective, error) {
	arguments, err := argumentsToPb(directive.Arguments)
	if err != nil {
		return nil, err
	}

	return managementpb.TypeDirective_builder{
		Name:      directive.Name,
		Arguments: arguments,
	}.Build(), nil
}

func argumentsToPb(arguments []schema.Argument) ([]*managementpb.TypeArgument, error) {
	var pbArguments []*managementpb.TypeArgument

	for _, argument := range arguments {
		pbArgument, err := argumentToPb(&argument)
		if err != nil {
			return nil, err
		}

		pbArguments = append(pbArguments, pbArgument)
	}

	return pbArguments, nil
}

func argumentToPb(argument *schema.Argument) (*managementpb.TypeArgument, error) {
	valueBytes, err := json.Marshal(argument.Value)
	if err != nil {
		return nil, err
	}

	return managementpb.TypeArgument_builder{
		Name:  argument.Name,
		Value: string(valueBytes),
	}.Build(), nil
}
