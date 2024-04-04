package api

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	openApiParser "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/parser"
	"path/filepath"
)

import "github.com/rocket-generator/rocket-generator-cli/modules/commands/create"

type Command struct {
}

func (c *Command) Execute(arguments Arguments) error {
	var err error
	typeMapper := arguments.TypeMapper
	if typeMapper == nil {
		typeMapperFilePath := filepath.Join(arguments.ProjectPath, "templates", "data", "types.json")
		typeMapper, err = data_mapper.Parse(typeMapperFilePath)
		if err != nil {
			return err
		}
	}

	api := arguments.ApiSpec
	if api == nil {
		api, err = openApiParser.Parse(
			arguments.ApiFileName,
			"app",
			"",
			"",
			typeMapper,
		)
		arguments.ApiSpec = api
	}

	payload := Payload{
		Type:        arguments.Type,
		ProjectPath: arguments.ProjectPath,
		Request:     arguments.Request,
		ApiSpec:     api,
		Debug:       arguments.Debug,
	}

	if payload.Request == nil {
		for _, request := range payload.ApiSpec.Requests {
			if request.Path == arguments.Path && request.Method.Original == arguments.Method {
				payload.Request = request
				break
			}
		}
	}

	err = create.GenerateFileFromTemplate(payload.ProjectPath, payload.Type, *payload.Request)
	if err != nil {
		return err
	}
	err = create.GenerateEmbeddedPartFromTemplate(payload.ProjectPath, payload.Type, *payload.Request)
	if err != nil {
		return err
	}
	return nil
}
