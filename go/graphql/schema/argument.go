package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/pkg/errors"
)

type Argument struct {
	Name  string
	Kind  string
	Value any
}

func (a *Argument) GetString() (string, error) {
	argString, err := getArgumentString(a.Name, a.Value)
	if err != nil {
		return "", errors.Wrap(err, "error marshaling schema argument")
	}
	return argString, nil
}

func (a *Argument) Equals(other *Argument) bool {
	if a.Name != other.Name {
		return false
	}

	return isEqual(a.Value, other.Value)
}

func isEqual(v1 any, v2 any) bool {
	return reflect.DeepEqual(v1, v2)
}

func getArgumentString(name string, value any) (string, error) {
	valueString, err := getValueString(value)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s: %s", name, valueString), nil
}

func getValueString(value any) (string, error) {
	if sliceData, err := getSliceData(value); err == nil {
		var stringElements []string

		for _, data := range sliceData {
			childString, err := getValueString(data)
			if err != nil {
				return "", err
			}

			stringElements = append(stringElements, childString)
		}

		return fmt.Sprintf("[%s]", strings.Join(stringElements, ",")), nil
	}

	mapValue, ok := value.(map[string]any)
	if ok {
		argString, err := getObjectArgumentMapString(mapValue)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("{ %s }", argString), nil
	}

	valueBytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(valueBytes), nil
}

func getObjectArgumentMapString(mapValue map[string]any) (string, error) {
	var buffer bytes.Buffer
	i := 0

	var keys []string

	for key := range mapValue {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		value := mapValue[key]

		if i != 0 {
			buffer.WriteString(", ")
		}

		argString, err := getArgumentString(key, value)
		if err != nil {
			return "", err
		}

		buffer.WriteString(argString)
		i++
	}

	return buffer.String(), nil
}

func getSliceData(data any) ([]any, error) {
	if data == nil {
		return []any{}, nil
	}

	if sliceData, ok := data.([]any); ok {
		return sliceData, nil
	}

	return nil, fmt.Errorf("expected array but got %T", data)
}
