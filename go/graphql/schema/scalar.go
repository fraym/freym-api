package schema

import (
	"fmt"
	"strings"
)

type Scalar struct {
	Name       string
	Directives []Directive
}

func (s *Scalar) GetDirective(name string) (*Directive, error) {
	for _, directive := range s.Directives {
		if directive.Name == name {
			return &directive, nil
		}
	}

	return nil, fmt.Errorf("could not find directive `%s` on scalart `%s`", name, s.Name)
}

func (s *Scalar) GetString() (string, error) {
	var directivesString strings.Builder

	for _, arg := range s.Directives {
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

	return fmt.Sprintf("scalar %s%s", s.Name, directivesString.String()), nil
}

func (s *Scalar) Equals(other *Scalar) bool {
	if s.Name != other.Name || len(s.Directives) != len(other.Directives) {
		return false
	}

	for _, directive := range s.Directives {
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
