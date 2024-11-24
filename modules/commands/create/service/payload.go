package service

import (
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type PayloadModel struct {
	Name        create.Name
	IsCRUDModel bool
}

// Payload ...
type Payload struct {
	Type            string
	Name            create.Name
	RelatedModels   []PayloadModel
	IsAuthService   bool
	ProjectPath     string
	RelatedResponse *objects.Response
	Debug           bool
}
