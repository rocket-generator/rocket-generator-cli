package payload

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec"
)

type Payload struct {
	TemplateName     string
	ProjectName      string
	ProjectBasePath  string
	ProjectPath      string
	ApiFileName      string
	DatabaseFileName string
	OrganizationName string
	OpenAPISpec      *openapispec.API
	DatabaseSchema   *databaseschema.Schema
	TypeMapper       *data_mapper.Mapper
	Count            int
	Debug            bool
}
