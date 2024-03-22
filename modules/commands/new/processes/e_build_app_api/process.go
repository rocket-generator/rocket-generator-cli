package e_build_app_api

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	createDtoCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/dto"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, request := range payload.OpenAPISpec.Requests {
		green := color.New(color.FgGreen)
		_, _ = green.Println("* Generate files from api: " + request.Method.Original + " " + request.Path)
		if payload.Debug {
			_byte, _ := json.MarshalIndent(request, "", "    ")
			fmt.Println(string(_byte))
		}
		if err := process.generateFileFromTemplate(*request, payload); err != nil {
			return nil, err
		}
		if !request.HasStatusResponse && !request.SuccessResponse.IsList {
			// Require DTO
			argument := createDtoCommand.Arguments{
				Type:              "dto",
				Name:              request.SuccessResponse.Schema.Name.Default.Title,
				RelatedModelNames: []string{request.TargetModel},
				RelatedResponse:   request.SuccessResponse,
				ProjectPath:       payload.ProjectPath,
				Debug:             payload.Debug,
			}
			command := createDtoCommand.Command{}
			err := command.Execute(argument)
			if err != nil {
				return nil, err
			}
		}
	}
	err := process.generateEmbeddedPartFromTemplate(payload.OpenAPISpec, payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
