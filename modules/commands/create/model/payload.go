package model

import (
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	databaseObject "github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
)

// Payload ...
type Payload struct {
	Type           string
	Name           create.Name
	DatabaseSchema *databaseObject.Schema
	Entity         *databaseObject.Entity
	TypeMapper     *data_mapper.Mapper
	ProjectPath    string
	Debug          bool
}
