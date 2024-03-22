package api

import "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
import "github.com/rocket-generator/rocket-generator-cli/modules/commands/create"

// Payload ...
type Payload struct {
	Type            string
	Name            create.Name
	RelatedModels   []create.Name
	ProjectPath     string
	RelatedResponse *objects.Response
	Debug           bool
}
