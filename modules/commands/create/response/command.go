package response

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	openApiParser "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/parser"
	"github.com/stoewer/go-strcase"

	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
)

type Command struct {
}

func (c *Command) Execute(arguments Arguments) error {
	var err error

	ignoreKey := strings.ToLower(arguments.Name)
	if arguments.IgnoreList != nil {
		if _, ok := arguments.IgnoreList.Responses[ignoreKey]; ok {
			return nil
		}
	}

	// Type mapper
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
	}

	// Loop api models to find response name
	responseName := strcase.SnakeCase(arguments.Name)
	// find schema by name and check ok
	schema, ok := api.Schemas[responseName]
	if !ok {
		return errors.New("schema " + responseName + " not found")
	}

	fmt.Println("Create response: " + responseName)
	payload := &Payload{
		Type:              "response",
		Name:              create.GenerateName(responseName),
		ProjectPath:       arguments.ProjectPath,
		Schema:            schema,
		ApiSpec:           api,
		Debug:             arguments.Debug,
		HasStatusResponse: arguments.HasStatusResponse,
		IsList:            arguments.IsList,
		IgnoreList:        arguments.IgnoreList,
	}

	if payload.Name.Singular.Snake == "status" {
		payload.HasStatusResponse = true
	} else {
		payload.HasStatusResponse = false
	}

	if payload.Name.Plural.Snake == payload.Name.Default.Snake {
		hasData := false
		hasCount := false
		for _, property := range payload.Schema.Properties {
			if property.Name.Default.Camel == "data" {
				hasData = true
			}
			if property.Name.Default.Camel == "count" {
				hasCount = true
			}
		}
		payload.IsList = hasData && hasCount
	} else {
		payload.IsList = false
	}

	err = create.GenerateFileFromTemplate(payload.ProjectPath, payload.Type, payload)
	if err != nil {
		return err
	}

	err = create.GenerateEmbeddedPartFromTemplate(payload.ProjectPath, payload.Type, payload)
	if err != nil {
		return err
	}

	for _, property := range schema.Properties {
		if property.Type == "object" {
			childSchemaName := property.Reference
			childSchema, ok := payload.ApiSpec.Schemas[childSchemaName]
			if ok {
				childArgument := Arguments{
					Type:        "response",
					Name:        childSchemaName,
					ApiFileName: arguments.ApiFileName,
					Schema:      childSchema,
					ApiSpec:     arguments.ApiSpec,
					TypeMapper:  arguments.TypeMapper,
					ProjectPath: arguments.ProjectPath,
					Debug:       arguments.Debug,
				}
				err = c.Execute(childArgument)
				if err != nil {
					return err
				}
			}
		} else if property.Type == "array" {
			childSchemaName := property.ArrayItemName
			childSchema, ok := payload.ApiSpec.Schemas[childSchemaName]
			if ok {
				childArgument := Arguments{
					Type:        "response",
					Name:        childSchemaName,
					ApiFileName: arguments.ApiFileName,
					Schema:      childSchema,
					ApiSpec:     arguments.ApiSpec,
					TypeMapper:  arguments.TypeMapper,
					ProjectPath: arguments.ProjectPath,
					Debug:       arguments.Debug,
				}
				err = c.Execute(childArgument)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
