package schema

import (
	"fmt"
	"strings"
)

type Schema struct {
	Objects []Object
	Enums   []Enum
	Scalars []Scalar
}

func (s *Schema) GetObject(name string) (*Object, error) {
	for _, object := range s.Objects {
		if object.Name == name {
			return &object, nil
		}
	}

	return nil, fmt.Errorf("object `%s` does not exist in given schema", name)
}

func (s *Schema) GetEnum(name string) (*Enum, error) {
	for _, enum := range s.Enums {
		if enum.Name == name {
			return &enum, nil
		}
	}

	return nil, fmt.Errorf("enum `%s` does not exist in given schema", name)
}

func (s *Schema) GetScalar(name string) (*Scalar, error) {
	for _, scalar := range s.Scalars {
		if scalar.Name == name {
			return &scalar, nil
		}
	}

	return nil, fmt.Errorf("scalar `%s` does not exist in given schema", name)
}

func (s *Schema) GetString() (string, error) {
	var builder strings.Builder

	for _, enum := range s.Enums {
		enumString, err := enum.GetString()
		if err != nil {
			return "", err
		}

		if _, err := builder.WriteString(fmt.Sprintf("%s\n", enumString)); err != nil {
			return "", err
		}
	}

	for _, object := range s.Objects {
		objectString, err := object.GetString()
		if err != nil {
			return "", err
		}

		if _, err := builder.WriteString(fmt.Sprintf("%s\n", objectString)); err != nil {
			return "", err
		}
	}

	for _, scalar := range s.Scalars {
		scalarString, err := scalar.GetString()
		if err != nil {
			return "", err
		}

		if _, err := builder.WriteString(fmt.Sprintf("%s\n", scalarString)); err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}

func (s *Schema) Equals(other *Schema) bool {
	if len(s.Objects) != len(other.Objects) || len(s.Enums) != len(other.Enums) ||
		len(s.Scalars) != len(other.Scalars) {
		return false
	}

	for _, object := range s.Objects {
		otherObject, err := other.GetObject(object.Name)
		if err != nil || !object.Equals(otherObject) {
			return false
		}
	}

	for _, enum := range s.Enums {
		otherEnum, err := other.GetEnum(enum.Name)
		if err != nil || !enum.Equals(otherEnum) {
			return false
		}
	}

	for _, scalar := range s.Scalars {
		scalarEnum, err := other.GetScalar(scalar.Name)
		if err != nil || !scalar.Equals(scalarEnum) {
			return false
		}
	}

	return true
}
