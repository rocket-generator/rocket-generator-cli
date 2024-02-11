package payload

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	databaseObject "github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	apiObjects "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type Payload struct {
	TemplateName     string
	ProjectName      string
	ProjectBasePath  string
	ProjectPath      string
	ApiFileName      string
	DatabaseFileName string
	OrganizationName string
	OpenAPISpec      *apiObjects.API
	DatabaseSchema   *databaseObject.Schema
	TypeMapper       *data_mapper.Mapper
	Count            int
	Debug            bool
	HasAdminAPI      bool
}
