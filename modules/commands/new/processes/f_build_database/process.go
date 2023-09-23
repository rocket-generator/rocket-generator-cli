package f_build_database

import (
	"encoding/json"
	"fmt"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, entity := range payload.DatabaseSchema.Entities {
		if payload.Debug {
			_byte, _ := json.MarshalIndent(entity, "", "    ")
			// _byte, _ := json.Marshal(request)
			fmt.Println(string(_byte))
		}
		if err := process.generateFileFromTemplate(*entity, payload); err != nil {
			return nil, err
		}
	}
	err := process.generateEmbeddedPartFromTemplate(payload.DatabaseSchema.Entities, payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
