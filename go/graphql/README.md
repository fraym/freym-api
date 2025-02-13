# graphql

This package provides helpers that simplify handling of graphql related things.

## Schema parsing

Create a struct representaiton of your schema objects:

```go
schema, err := graphql.NewSchema(schemaString)
if err != nil {
	panic(err)
}
```

The returned schema struct offers useful functions that can be used to inspect the schema.

## Schema validation

You can use this library to validate a given graphql schema against your requirements:

### Instantiationg the validators

There are validators for directives, fields and objeccts.

```go
validator := graphql.NewSchemaValidator(&validator.ObjectConfig{
	// configure your object requirements here
}, &validator.DirectiveConfig{
	// configure your directive requirements here
}, &validator.FieldConfig{
	// configure your field requirements here
})

if err := validator.Validate(schema); err != nil {
	panic(err)
}
```

## Schema printing

You can print your schema to a graphql schema string using:

```go
schema, err := graphql.NewSchema(schemaString)
if err != nil {
	panic(err)
}

generatedSchemaString, err := schema.GetString()
if err != nil {
	panic(err)
}
```

The `generatedSchemaString` will be logically equal to the schama that you one provided as `schemaString` (It might have differences in formatting, therefore it will not be 100% equal).

## Testing

For testing purposes the schema package provides mocks that can be used to mock the validators in your project:

ObjectValidator:

```go
validator := validator.MockObjectValidator{}
```

FieldValidator:

```go
validator := validator.MockFieldValidator{}
```

DirectiveValidator:

```go
validator := validator.MockDirectiveValidator{}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
