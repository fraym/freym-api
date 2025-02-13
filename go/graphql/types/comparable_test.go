package types_test

import (
	"testing"

	"github.com/fraym/freym-api/go/graphql/types"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/stretchr/testify/assert"
)

func TestGetComparableType_Empty(t *testing.T) {
	_, err := types.GetComparable(nil)
	assert.Error(t, err)
}

func TestGetComparableType_Named(t *testing.T) {
	expectedValue := "test-name"

	astType := ast.NewNamed(&ast.Named{
		Name: ast.NewName(&ast.Name{
			Value: expectedValue,
		}),
	})

	types, err := types.GetComparable(astType)
	assert.NoError(t, err)
	assert.True(t, types.Equals([]string{kinds.Named, expectedValue}))
}

func TestGetComparableType_List(t *testing.T) {
	expectedValue := "test-name"

	astType := ast.NewList(&ast.List{
		Type: ast.NewNamed(&ast.Named{
			Name: ast.NewName(&ast.Name{
				Value: expectedValue,
			}),
		}),
	})

	types, err := types.GetComparable(astType)
	assert.NoError(t, err)
	assert.True(t, types.Equals([]string{kinds.List, kinds.Named, expectedValue}))
}

func TestGetComparableType_NonNull(t *testing.T) {
	expectedValue := "test-name"

	astType := ast.NewNonNull(&ast.NonNull{
		Type: ast.NewNamed(&ast.Named{
			Name: ast.NewName(&ast.Name{
				Value: expectedValue,
			}),
		}),
	})

	types, err := types.GetComparable(astType)
	assert.NoError(t, err)
	assert.True(t, types.Equals([]string{kinds.NonNull, kinds.Named, expectedValue}))
}

func TestComparable_GetString(t *testing.T) {
	expectedString := "[String!]!"

	c := types.Comparable{kinds.NonNull, kinds.List, kinds.NonNull, kinds.Named, "String"}

	str, err := c.GetString()
	assert.NoError(t, err)
	assert.Equal(t, expectedString, str)
}
