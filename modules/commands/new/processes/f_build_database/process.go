package f_build_database

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	createModelCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/model"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, entity := range payload.DatabaseSchema.Entities {
		green := color.New(color.FgGreen)
		_, _ = green.Println("* Generate files from db table: " + entity.Name.Original)
		if payload.Debug {
			_byte, _ := json.MarshalIndent(entity, "", "    ")
			fmt.Println(string(_byte))
		}

		argument := createModelCommand.Arguments{
			Type:           "model",
			ProjectPath:    payload.ProjectPath,
			Name:           entity.Name.Original,
			DatabaseSchema: payload.DatabaseSchema,
			Entity:         entity,
			TypeMapper:     payload.TypeMapper,
			Debug:          payload.Debug,
		}
		command := createModelCommand.Command{}
		err := command.Execute(argument)
		if err != nil {
			return nil, err
		}
	}
	return payload, nil
}
