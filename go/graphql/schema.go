package graphql

import (
	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/fraym/freym-api/go/graphql/schema"
)

func NewSchema(schema string) (*schema.Schema, error) {
	return parser.GetSchemaFromString(schema)
}
