package projections

import (
	"github.com/Becklyn/go-wire-core/json"
	"github.com/fraym/freym-api/go/proto/projections/managementpb"
)

type DeploymentOptions struct {
	DangerouslyRemoveGdprFields bool
	Target                      managementpb.DeploymentTarget
	Force                       bool
}

func (o *DeploymentOptions) toPb() *managementpb.DeploymentOptions {
	if o == nil {
		return nil
	}

	return managementpb.DeploymentOptions_builder{
		DangerouslyRemoveGdprFields: o.DangerouslyRemoveGdprFields,
		Target:                      o.Target,
		Force:                       o.Force,
	}.Build()
}

type View struct {
	Name       string
	Sql        string
	Directives []TypeDirective
	Fields     []TypeField
}

func viewListToPb(list []View) ([]*managementpb.View, error) {
	var newList []*managementpb.View

	for _, view := range list {
		newEnum, err := view.toPb()
		if err != nil {
			return nil, err
		}

		newList = append(newList, newEnum)
	}

	return newList, nil
}

func (v *View) toPb() (*managementpb.View, error) {
	directives, err := typeDirectiveListToPb(v.Directives)
	if err != nil {
		return nil, err
	}

	fields, err := typeFieldListToPb(v.Fields)
	if err != nil {
		return nil, err
	}

	return managementpb.View_builder{
		Name:       v.Name,
		Sql:        v.Sql,
		Directives: directives,
		Fields:     fields,
	}.Build(), nil
}

type ObjectType struct {
	Name       string
	Directives []TypeDirective
	Fields     []TypeField
}

func objectTypeListToPb(list []ObjectType) ([]*managementpb.ObjectType, error) {
	var newList []*managementpb.ObjectType

	for _, nestedType := range list {
		newType, err := nestedType.toPb()
		if err != nil {
			return nil, err
		}

		newList = append(newList, newType)
	}

	return newList, nil
}

func (c *ObjectType) toPb() (*managementpb.ObjectType, error) {
	directives, err := typeDirectiveListToPb(c.Directives)
	if err != nil {
		return nil, err
	}

	fields, err := typeFieldListToPb(c.Fields)
	if err != nil {
		return nil, err
	}

	return managementpb.ObjectType_builder{
		Name:       c.Name,
		Directives: directives,
		Fields:     fields,
	}.Build(), nil
}

type EnumType struct {
	Name   string
	Values []string
}

func enumTypeListToPb(list []EnumType) ([]*managementpb.EnumType, error) {
	var newList []*managementpb.EnumType

	for _, enum := range list {
		newEnum, err := enum.toPb()
		if err != nil {
			return nil, err
		}

		newList = append(newList, newEnum)
	}

	return newList, nil
}

func (e *EnumType) toPb() (*managementpb.EnumType, error) {
	return managementpb.EnumType_builder{
		Name:   e.Name,
		Values: e.Values,
	}.Build(), nil
}

type TypeField struct {
	Name       string
	Type       []string
	Directives []TypeDirective
}

func typeFieldListToPb(list []TypeField) ([]*managementpb.TypeField, error) {
	var newList []*managementpb.TypeField

	for _, field := range list {
		newField, err := field.toPb()
		if err != nil {
			return nil, err
		}

		newList = append(newList, newField)
	}

	return newList, nil
}

func (f *TypeField) toPb() (*managementpb.TypeField, error) {
	directives, err := typeDirectiveListToPb(f.Directives)
	if err != nil {
		return nil, err
	}

	return managementpb.TypeField_builder{
		Name:       f.Name,
		Type:       f.Type,
		Directives: directives,
	}.Build(), nil
}

type TypeDirective struct {
	Name      string
	Arguments []TypeArgument
}

func typeDirectiveListToPb(list []TypeDirective) ([]*managementpb.TypeDirective, error) {
	var newList []*managementpb.TypeDirective

	for _, directive := range list {
		newDirective, err := directive.toPb()
		if err != nil {
			return nil, err
		}

		newList = append(newList, newDirective)
	}

	return newList, nil
}

func (d *TypeDirective) toPb() (*managementpb.TypeDirective, error) {
	arguments, err := typeArgumentListToPb(d.Arguments)
	if err != nil {
		return nil, err
	}

	return managementpb.TypeDirective_builder{
		Name:      d.Name,
		Arguments: arguments,
	}.Build(), nil
}

type TypeArgument struct {
	Name  string
	Value any
}

func typeArgumentListToPb(list []TypeArgument) ([]*managementpb.TypeArgument, error) {
	var newList []*managementpb.TypeArgument

	for _, argument := range list {
		newArgument, err := argument.toPb()
		if err != nil {
			return nil, err
		}

		newList = append(newList, newArgument)
	}

	return newList, nil
}

func (a *TypeArgument) toPb() (*managementpb.TypeArgument, error) {
	bytes, err := json.Marshal(a.Value)
	if err != nil {
		return nil, err
	}

	return managementpb.TypeArgument_builder{
		Name:  a.Name,
		Value: string(bytes),
	}.Build(), nil
}
