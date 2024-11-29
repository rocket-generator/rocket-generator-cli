package dto

import (
	databaseObject "github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type Arguments struct {
	Type              string
	Name              string
	RelatedModelNames []string
	RelatedMainModel  *databaseObject.Entity
	RelatedResponse   *objects.Response
	ProjectPath       string
	Debug             bool
}
