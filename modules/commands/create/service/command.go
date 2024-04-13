package service

import (
	"github.com/rocket-generator/rocket-generator-cli/internal/utilities"
	"github.com/stoewer/go-strcase"
)

import "github.com/rocket-generator/rocket-generator-cli/modules/commands/create"

type Command struct {
}

func (c *Command) Execute(arguments Arguments) error {
	payload := Payload{
		Type:            arguments.Type,
		Name:            create.GenerateName(utilities.RemovePostfix(arguments.Name, strcase.UpperCamelCase(arguments.Type))),
		RelatedModels:   []create.Name{},
		IsAuthService:   arguments.IsAuthService,
		RelatedResponse: arguments.RelatedResponse,
		ProjectPath:     arguments.ProjectPath,
		Debug:           arguments.Debug,
	}

	for _, relatedModelName := range arguments.RelatedModelNames {
		payload.RelatedModels = append(payload.RelatedModels, create.GenerateName(relatedModelName))
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
