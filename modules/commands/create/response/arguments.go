package response

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/ignore_list"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type Arguments struct {
	Type              string
	Name              string
	ApiFileName       string
	Schema            *objects.Schema
	ApiSpec           *objects.API
	TypeMapper        *data_mapper.Mapper
	IgnoreList        *ignore_list.IgnoreList
	ProjectPath       string
	HasStatusResponse bool
	IsList            bool
	Debug             bool
}
