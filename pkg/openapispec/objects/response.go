package objects

import (
	databaseObject "github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
)

// Response ...
type Response struct {
	StatusCode    string
	Success       bool
	Schema        *Schema
	RelatedEntity *databaseObject.Entity
	IsList        bool
}
