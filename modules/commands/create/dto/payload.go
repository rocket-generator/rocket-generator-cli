package dto

import (
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	databaseObject "github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

// Payload ...
type Payload struct {
	Type             string
	Name             create.Name
	RelatedMainModel *databaseObject.Entity
	RelatedModels    []create.Name
	ProjectPath      string
	RelatedResponse  *objects.Response
	Debug            bool
}
