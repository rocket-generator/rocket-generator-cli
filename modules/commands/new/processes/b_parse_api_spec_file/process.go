package b_parse_api_spec_file

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	openApiParser "github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/parser"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	api, err := openApiParser.Parse(
		payload.ApiFileName,
		"app",
		payload.ProjectName,
		payload.OrganizationName,
		payload.TypeMapper,
	)
	if err != nil {
		return nil, err
	}
	payload.OpenAPISpec = api
	return process.GuessCRUD(payload)
}

func (process *Process) GuessCRUD(payload *newCommand.Payload) (*newCommand.Payload, error) {
	api := payload.OpenAPISpec
	if api == nil {
		return payload, nil
	}

	if len(api.Requests) == 0 {
		return payload, nil
	}

	for _, request := range api.Requests {
		if request.SuccessResponse.Schema.Name.Default.Title == "Status" {
			request.TargetModel = &request.PathLastElement
		} else {
			request.TargetModel = &request.SuccessResponse.Schema.Name
		}
		switch request.Method.Title {
		case "Delete":
			request.RequestType = "crud"
			request.RequestSubType = "delete"
		case "Put":
			request.RequestType = "crud"
			request.RequestSubType = "update"
		case "Post":
			if request.PathName.Plural.Snake == request.PathName.Default.Snake {
				request.RequestType = "crud"
				request.RequestSubType = "create"
			}
		case "Get":
			if request.PathName.Plural.Snake == request.PathName.Default.Snake {
				request.RequestType = "crud"
				request.RequestSubType = "index"
			} else if request.PathName.Singular.Snake == request.PathName.Default.Snake {
				request.RequestType = "crud"
				request.RequestSubType = "get"
			}
		}
	}

	return payload, nil
}
