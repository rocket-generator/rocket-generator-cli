package d_build_app_api

import (
	"fmt"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	for _, request := range payload.OpenAPISpec.Requests {
		_byte, _ := json.MarshalIndent(request, "", "    ")
		// _byte, _ := json.Marshal(request)
		fmt.Println(string(_byte))
		if err := process.generateFileFromTemplate(*request, payload); err != nil {
			return nil, err
		}
	}
	return payload, nil
}
