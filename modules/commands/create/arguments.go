package create

import "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"

type Arguments struct {
	Type              string
	Name              string
	RelatedModelNames []string
	RelatedResponse   *objects.Response
	ProjectPath       string
	Debug             bool
}
