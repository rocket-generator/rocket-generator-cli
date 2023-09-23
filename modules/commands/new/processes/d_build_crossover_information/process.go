package d_build_crossover_information

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, request := range payload.OpenAPISpec.Requests {
		successResponse := request.SuccessResponse
		if successResponse != nil {
			relatedEntity, isList := findSameNameDatabaseEntity(successResponse.Schema.Name.Default.Title, payload)
			if relatedEntity != nil {
				request.SuccessResponse.RelatedEntity = relatedEntity
				request.SuccessResponse.IsList = isList
			}
		}
	}
	return payload, nil
}

func findSameNameDatabaseEntity(entityName string, payload *newCommand.Payload) (*objects.Entity, bool) {
	for _, entity := range payload.DatabaseSchema.Entities {
		if entity.Name.Singular.Title == entityName {
			return entity, false
		}
		if entity.Name.Plural.Title == entityName {
			return entity, true
		}
	}
	return nil, false
}
