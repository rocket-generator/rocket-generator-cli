package model

import (
	"errors"
	"github.com/rocket-generator/rocket-generator-cli/internal/utilities"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/parser"
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

	schema := arguments.DatabaseSchema
	if schema == nil {
		schema, err = parser.ParseDBML(
			arguments.DatabaseFileName,
			"",
			"",
			typeMapper,
		)
		if err != nil {
			return err
		}
	}

	payload := Payload{
		Type:            arguments.Type,
		Name:            create.GenerateName(utilities.RemovePostfix(arguments.Name, strcase.UpperCamelCase(arguments.Type))),
		DatabaseSchema:  schema,
		Entity:          arguments.Entity,
		TypeMapper:      typeMapper,
		ProjectPath:     arguments.ProjectPath,
		Debug:           arguments.Debug,
		Authenticatable: arguments.Authenticatable,
	}

	payload.Entity = arguments.Entity
	if payload.Entity == nil {
		for _, entity := range payload.DatabaseSchema.Entities {
			if entity.Name.Original == arguments.Name {
				payload.Entity = entity
				break
			}
		}
	}
	if payload.Entity == nil {
		return errors.New("entity not found: " + arguments.Name)
	}

	err = create.GenerateFileFromTemplate(payload.ProjectPath, payload.Type, payload.Entity)
	if err != nil {
		return err
	}

	err = create.GenerateEmbeddedPartFromTemplate(payload.ProjectPath, payload.Type, payload.Entity)
	if err != nil {
		return err
	}
	return nil
}
