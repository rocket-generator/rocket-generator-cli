package b_parse_api_info_file

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/parser"
)

type Process struct {
}

type API struct {
	Path           string          `json:"path"`
	Method         string          `json:"method"`
	Type           string          `json:"type"`
	SubType        string          `json:"subType"`
	Group          string          `json:"group"`
	RequireAuth    bool            `json:"requireAuth"`
	RequiredRoles  []string        `json:"requiredRoles"`
	TargetModel    string          `json:"targetModel"`
	AncestorModels []AncestorModel `json:"ancestorModels"`
}

type AncestorModel struct {
	Name      string `json:"name"`
	Parameter string `json:"parameter"`
	Column    string `json:"column"`
}

type APIs []API

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	if payload.ApiInfoFileName == "" {
		return payload, nil
	}
	err := parser.ParseAPIInfoFile(payload.ApiInfoFileName, payload.OpenAPISpec)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
