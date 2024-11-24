package service

import (
	"github.com/rocket-generator/rocket-generator-cli/internal/utilities"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	"github.com/stoewer/go-strcase"
)

type Command struct {
}

func (c *Command) Execute(arguments Arguments) error {
	payload := Payload{
		Type:            arguments.Type,
		Name:            create.GenerateName(utilities.RemovePostfix(arguments.Name, strcase.UpperCamelCase(arguments.Type))),
		RelatedModels:   []PayloadModel{},
		IsAuthService:   arguments.IsAuthService,
		RelatedResponse: arguments.RelatedResponse,
		ProjectPath:     arguments.ProjectPath,
		Debug:           arguments.Debug,
	}

	for _, relatedModelName := range arguments.RelatedModelNames {
		payload.RelatedModels = append(payload.RelatedModels, PayloadModel{
			Name:        create.GenerateName(relatedModelName),
			IsCRUDModel: false,
		})
	}

	for _, relatedModelWithCRUDName := range arguments.RelatedModelWithCRUDNames {
		payload.RelatedModels = append(payload.RelatedModels, PayloadModel{
			Name:        create.GenerateName(relatedModelWithCRUDName),
			IsCRUDModel: true,
		})
	}

	err := create.GenerateFileFromTemplate(payload.ProjectPath, payload.Type, payload)
	if err != nil {
		return err
	}
	err = create.GenerateEmbeddedPartFromTemplate(payload.ProjectPath, payload.Type, payload)
	if err != nil {
		return err
	}
	return nil
}
