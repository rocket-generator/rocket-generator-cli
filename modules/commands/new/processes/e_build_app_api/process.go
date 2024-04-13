package e_build_app_api

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	createApiCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/api"
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
		apiArgument := createApiCommand.Arguments{
			Type:            "api",
			ProjectPath:     payload.ProjectPath,
			Method:          request.Method.Original,
			Path:            request.Path,
			TypeMapper:      payload.TypeMapper,
			ApiSpec:         payload.OpenAPISpec,
			ApiFileName:     payload.ApiFileName,
			ApiInfoFileName: payload.ApiInfoFileName,
			Request:         request,
			Debug:           payload.Debug,
		}
		apiCommand := createApiCommand.Command{}
		err := apiCommand.Execute(apiArgument)
		if err != nil {
			return nil, err
		}
		if !request.HasStatusResponse && !request.SuccessResponse.IsList {
			// Require DTO
			dtoArgument := createDtoCommand.Arguments{
				Type:              "dto",
				Name:              request.SuccessResponse.Schema.Name.Default.Title,
				RelatedModelNames: nil,
				RelatedResponse:   request.SuccessResponse,
				ProjectPath:       payload.ProjectPath,
				Debug:             payload.Debug,
			}
			dtoCommand := createDtoCommand.Command{}
			err := dtoCommand.Execute(dtoArgument)
			if err != nil {
				return nil, err
			}
		}
	}
	return payload, nil
}
