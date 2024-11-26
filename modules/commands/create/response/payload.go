package response

import (
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/ignore_list"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

// Payload ...
type Payload struct {
	Type              string
	Name              create.Name
	ProjectPath       string
	IgnoreList        *ignore_list.IgnoreList
	TypeMapper        *data_mapper.Mapper
	Schema            *objects.Schema
	HasStatusResponse bool
	IsList            bool
	ApiSpec           *objects.API
	Debug             bool
}
