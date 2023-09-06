package payload

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	objects2 "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type Payload struct {
	TemplateName     string
	ProjectName      string
	ProjectBasePath  string
	ProjectPath      string
	ApiFileName      string
	DatabaseFileName string
	OrganizationName string
	OpenAPISpec      *objects2.API
	DatabaseSchema   *objects.Schema
	TypeMapper       *data_mapper.Mapper
	Count            int
	Debug            bool
}
