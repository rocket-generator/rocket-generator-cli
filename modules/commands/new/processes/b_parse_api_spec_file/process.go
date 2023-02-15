package b_parse_api_spec_file

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	openApiParser "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/parser"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	api, err := openApiParser.Parse(payload.ApiFileName, "app", payload.ProjectName, payload.OrganizationName)
	if err != nil {
		return nil, err
	}
	payload.OpenAPISpec = api
	return payload, nil
}
