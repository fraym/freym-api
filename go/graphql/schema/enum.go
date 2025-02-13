package schema

import (
	"fmt"
	"strings"
)

type Enum struct {
	Name   string
	Values []string
}

func (e *Enum) GetString() (string, error) {
	var valuesBuilder strings.Builder

	for _, value := range e.Values {
		if _, err := valuesBuilder.WriteString(fmt.Sprintf("\t%s\n", value)); err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("enum %s {\n%s}", e.Name, valuesBuilder.String()), nil
}

func (e *Enum) Equals(other *Enum) bool {
	if e.Name != other.Name || len(e.Values) != len(other.Values) {
		return false
	}

	for i, value := range e.Values {
		if value != other.Values[i] {
			return false
		}
	}

	return true
}
