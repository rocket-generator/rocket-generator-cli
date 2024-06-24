package api

import (
	"errors"
	"fmt"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	openApiParser "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/parser"
	"github.com/stoewer/go-strcase"
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
		if err != nil {
			return err
		}
		arguments.ApiSpec = api

		if arguments.ApiInfoFileName != "" {
			err := openApiParser.ParseAPIInfoFile(arguments.ApiInfoFileName, arguments.ApiSpec)
			if err != nil {
				return err
			}
		}

	}

	payload := Payload{
		Type:        arguments.Type,
		ProjectPath: arguments.ProjectPath,
		Request:     arguments.Request,
		ApiSpec:     api,
		Debug:       arguments.Debug,
	}

	method := strcase.LowerCamelCase(arguments.Method)
	if payload.Request == nil {
		for _, request := range payload.ApiSpec.Requests {
			fmt.Println(request.Path, arguments.Path, request.Method.Original, method)
			if request.Path == arguments.Path && request.Method.Snake == method {
				payload.Request = request
				break
			}
		}
	}

	if payload.Request == nil {
		return errors.New("request not found")
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
