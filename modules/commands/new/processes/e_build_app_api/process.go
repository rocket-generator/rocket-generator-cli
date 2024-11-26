package e_build_app_api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	createApiCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/api"
	createDtoCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/dto"
	createResponseCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/response"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, request := range payload.OpenAPISpec.Requests {
		ignoreKey := strings.ToLower(request.Method.Camel) + " " + strings.ToLower(request.Path)
		// Ignore if ignoreKey is in payload.IgnoreList
		if _, ok := payload.IgnoreList.Endpoints[ignoreKey]; ok {
			yellow := color.New(color.FgYellow)
			_, _ = yellow.Println("* Ignore api: " + ignoreKey)
			continue
		}
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
		err = process.createResponseRecursively(payload, request.SuccessResponse.Schema)
		if err != nil {
			return nil, err
		}
		if !request.HasStatusResponse && !request.SuccessResponse.IsList {
			// Require DTO
			var targetModelNames []string
			if request.TargetModel != nil {
				targetModelNames = append(targetModelNames, request.TargetModel.Original)
			}
			dtoArgument := createDtoCommand.Arguments{
				Type:              "dto",
				Name:              request.SuccessResponse.Schema.Name.Default.Title,
				RelatedModelNames: targetModelNames,
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

func (process *Process) createResponseRecursively(payload *newCommand.Payload, schema *objects.Schema) error {
	responseArgument := createResponseCommand.Arguments{
		Type:        "response",
		Name:        schema.Name.Default.Snake,
		ApiFileName: payload.ApiFileName,
		Schema:      schema,
		ApiSpec:     payload.OpenAPISpec,
		TypeMapper:  payload.TypeMapper,
		IgnoreList:  payload.IgnoreList,
		ProjectPath: payload.ProjectPath,
		Debug:       payload.Debug,
	}
	responseCommand := createResponseCommand.Command{}
	err := responseCommand.Execute(responseArgument)
	if err != nil {
		return err
	}
	return nil
}
