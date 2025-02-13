package schema

import (
	"fmt"
	"strings"
)

type Object struct {
	Name       string
	Directives []Directive
	Fields     []Field
}

func (o *Object) GetDirective(name string) (*Directive, error) {
	for _, directive := range o.Directives {
		if directive.Name == name {
			return &directive, nil
		}
	}

	return nil, fmt.Errorf("could not find directive `%s` on object `%s`", name, o.Name)
}

func (o *Object) GetField(name string) (*Field, error) {
	for _, field := range o.Fields {
		if field.Name == name {
			return &field, nil
		}
	}

	return nil, fmt.Errorf("could not find field `%s` on object `%s`", name, o.Name)
}

func (o *Object) GetString() (string, error) {
	var (
		directivesBuilder strings.Builder
		fieldsBuilder     strings.Builder
	)

	for _, directive := range o.Directives {
		if _, err := directivesBuilder.WriteRune(' '); err != nil {
			return "", err
		}

		directiveString, err := directive.GetString()
		if err != nil {
			return "", err
		}

		if _, err := directivesBuilder.WriteString(directiveString); err != nil {
			return "", err
		}
	}

	for _, field := range o.Fields {
		fieldString, err := field.GetString()
		if err != nil {
			return "", err
		}

		if _, err := fieldsBuilder.WriteString(fieldString); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("type %s%s {\n%s}", o.Name, directivesBuilder.String(), fieldsBuilder.String()), nil
}

func (o *Object) Equals(other *Object) bool {
	if o.Name != other.Name || len(o.Directives) != len(other.Directives) || len(o.Fields) != len(other.Fields) {
		return false
	}

	for _, directive := range o.Directives {
		hasMatchingDirective := false

		for _, otherDirective := range other.Directives {
			if directive.Equals(&otherDirective) {
				hasMatchingDirective = true
				break
			}
		}

		if !hasMatchingDirective {
			return false
		}
	}

	for _, field := range o.Fields {
		hasMatchingField := false

		for _, otherField := range other.Fields {
			if field.Equals(&otherField) {
				hasMatchingField = true
				break
			}
		}

		if !hasMatchingField {
			return false
		}
	}

	return true
}
