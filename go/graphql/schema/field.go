package schema

import (
	"fmt"
	"strings"

	"github.com/fraym/freym-api/go/graphql/types"
)

type Field struct {
	Name       string
	Type       types.Comparable
	Directives []Directive
}

func (f *Field) GetDirective(name string) (*Directive, error) {
	for _, directive := range f.Directives {
		if directive.Name == name {
			return &directive, nil
		}
	}

	return nil, fmt.Errorf("could not find directive `%s` on field `%s`", name, f.Name)
}

func (f *Field) GetString() (string, error) {
	var directivesString strings.Builder

	for _, arg := range f.Directives {
		directiveString, err := arg.GetString()
		if err != nil {
			return "", err
		}
		if _, err := directivesString.WriteRune(' '); err != nil {
			return "", err
		}
		if _, err := directivesString.WriteString(directiveString); err != nil {
			return "", err
		}
	}

	typeString, err := f.Type.GetString()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("\t%s: %s%s\n", f.Name, typeString, directivesString.String()), nil
}

func (f *Field) Equals(other *Field) bool {
	if f.Name != other.Name || len(f.Directives) != len(other.Directives) || !f.Type.Equals(other.Type) {
		return false
	}

	for _, directive := range f.Directives {
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

	return true
}
