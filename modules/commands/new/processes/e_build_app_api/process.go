package e_build_app_api

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
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
	}
	err := process.generateEmbeddedPartFromTemplate(payload.OpenAPISpec.Requests, payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
