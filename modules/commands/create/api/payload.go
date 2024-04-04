package api

import "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
import "github.com/rocket-generator/rocket-generator-cli/modules/commands/create"

// Payload ...
type Payload struct {
	Type        string
	Name        create.Name
	ProjectPath string
	Request     *objects.Request
	ApiSpec     *objects.API
	Debug       bool
}
