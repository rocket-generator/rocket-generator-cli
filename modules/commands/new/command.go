package new

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/a_download_template"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/b_parse_api_spec_file"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/c_parse_database_spec_file"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/d_build_app_api"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/e_build_database"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/f_build_admin_api"
)

type Command struct {
}

func (c *Command) Execute(arguments Arguments) error {
	_payload := &payload.Payload{
		TemplateName:     arguments.Template,
		ProjectName:      arguments.ProjectName,
		ProjectBasePath:  arguments.ProjectBasePath,
		ProjectPath:      arguments.ProjectBasePath,
		ApiFileName:      arguments.ApiFileName,
		DatabaseFileName: arguments.DatabaseFileName,
		OrganizationName: arguments.OrganizationName,
		OpenAPISpec:      nil,
		DatabaseSchema:   nil,
		TypeMapper:       nil,
		Count:            0,
	}

	_processes := []processes.ProcessInterface{
		&a_download_template.Process{},
		&b_parse_api_spec_file.Process{},
		&c_parse_database_spec_file.Process{},
		&d_build_app_api.Process{},
		&e_build_database.Process{},
		&f_build_admin_api.Process{},
	}

	for _, process := range _processes {
		_payload, err := process.Execute(_payload)
		if err != nil {
			red := color.New(color.FgRed)
			_, _ = red.Println(err.Error())
			panic(err)
			return nil
		}
		_payload.Count++
	}

	fmt.Println(_payload.ProjectName + ": Project created successfully")

	return nil
}
