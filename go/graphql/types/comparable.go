package types

import (
	"fmt"

	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

type Comparable []string

func (t Comparable) Equals(other Comparable) bool {
	if len(t) != len(other) {
		return false
	}

	for i := range t {
		if t[i] != other[i] {
			return false
		}
	}

	return true
}

func GetComparable(t ast.Type, prev ...string) (Comparable, error) {
	if t == nil {
		return nil, fmt.Errorf("cannot determine type of empty type")
	}

	switch t.GetKind() {
	case kinds.List:
		subType := t.(*ast.List)
		return GetComparable(subType.Type, append(prev, kinds.List)...)
	case kinds.NonNull:
		subType := t.(*ast.NonNull)
		return GetComparable(subType.Type, append(prev, kinds.NonNull)...)
	case kinds.Named:
		subType := t.(*ast.Named)
		return append(prev, kinds.Named, subType.Name.Value), nil
	default:
		return nil, fmt.Errorf("unknown type `%s`", t.GetKind())
	}
}

func (t Comparable) GetString() (string, error) {
	if len(t) == 0 {
		return "", fmt.Errorf("cannot stringify invalid types.Comparable")
	}

	if len(t) == 1 {
		return t[0], nil
	}

	if t[0] == kinds.NonNull {
		str, err := t[1:].GetString()
		return fmt.Sprintf("%s!", str), err
	}

	if t[0] == kinds.List {
		str, err := t[1:].GetString()
		return fmt.Sprintf("[%s]", str), err
	}

	if t[0] == kinds.Named {
		return t[1:].GetString()
	}

	return "", fmt.Errorf("cannot stringify unknown types.Comparable `%s`", t[0])
}
