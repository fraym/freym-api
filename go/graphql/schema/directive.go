package schema

import (
	"fmt"
	"strings"
)

type Directive struct {
	Name      string
	Arguments []Argument
}

func (d *Directive) GetArgument(name string) (*Argument, error) {
	for _, arg := range d.Arguments {
		if arg.Name == name {
			return &arg, nil
		}
	}

	return nil, fmt.Errorf("could not find argument `%s` on directive `%s`", name, d.Name)
}

func (d *Directive) Equals(other *Directive) bool {
	if d.Name != other.Name || len(d.Arguments) != len(other.Arguments) {
		return false
	}

	for _, arg := range d.Arguments {
		otherArg, err := other.GetArgument(arg.Name)
		if err != nil || !arg.Equals(otherArg) {
			return false
		}
	}

	return true
}

func (d *Directive) GetString() (string, error) {
	if len(d.Arguments) == 0 {
		return fmt.Sprintf("@%s", d.Name), nil
	}

	var argsString strings.Builder

	max := len(d.Arguments) - 1
	for i, arg := range d.Arguments {
		argString, err := arg.GetString()
		if err != nil {
			return "", err
		}
		if _, err := argsString.WriteString(argString); err != nil {
			return "", err
		}

		if i < max {
			if _, err := argsString.WriteRune(','); err != nil {
				return "", err
			}
		}
	}

	return fmt.Sprintf("@%s(%s)", d.Name, argsString.String()), nil
}
