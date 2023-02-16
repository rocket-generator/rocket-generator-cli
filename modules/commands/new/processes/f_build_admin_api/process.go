package f_build_admin_api

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, entity := range payload.DatabaseSchema.Entities {

		if err := process.generateFileFromTemplate(*entity, payload); err != nil {
			return nil, err
		}
	}

	return payload, nil
}
