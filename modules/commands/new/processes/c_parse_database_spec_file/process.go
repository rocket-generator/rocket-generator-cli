package c_parse_database_spec_file

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/parser"
)

type Process struct {
}

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	schema, err := parser.ParseDBML(
		payload.DatabaseFileName,
		payload.ProjectName,
		payload.OrganizationName,
		payload.TypeMapper,
	)
	if err != nil {
		return nil, err
	}
	payload.DatabaseSchema = schema
	return payload, nil
}
