package f_build_database

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	createModelCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/model"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, entity := range payload.DatabaseSchema.Entities {
		ignoreKey := strings.ToLower(entity.Name.Original)
		// Ignore if ignoreKey is in payload.IgnoreList
		if payload.IgnoreList != nil {
			if _, ok := payload.IgnoreList.Tables[ignoreKey]; ok {
				yellow := color.New(color.FgYellow)
				_, _ = yellow.Println("* Ignore db table: " + ignoreKey)
				continue
			}
		}
		green := color.New(color.FgGreen)
		_, _ = green.Println("* Generate files from db table: " + entity.Name.Original)
		if payload.Debug {
			_byte, _ := json.MarshalIndent(entity, "", "    ")
			fmt.Println(string(_byte))
		}

		argument := createModelCommand.Arguments{
			Type:            "model",
			ProjectPath:     payload.ProjectPath,
			Name:            entity.Name.Original,
			DatabaseSchema:  payload.DatabaseSchema,
			Entity:          entity,
			TypeMapper:      payload.TypeMapper,
			Debug:           payload.Debug,
			Authenticatable: entity.Authenticatable,
		}
		command := createModelCommand.Command{}
		err := command.Execute(argument)
		if err != nil {
			return nil, err
		}
	}
	return payload, nil
}
