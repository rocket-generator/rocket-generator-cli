package new

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/a_download_template"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/b_parse_api_info_file"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/b_parse_api_spec_file"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/b_parse_service_definition_file"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/c_parse_database_info_file"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/c_parse_database_spec_file"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/d_build_crossover_information"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/e_build_app_api"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/f_build_database"
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/new/processes/g_build_admin_api"
)

type Command struct {
}

func (c *Command) Execute(arguments Arguments) error {
	_payload := &payload.Payload{
		TemplateName:         arguments.Template,
		ProjectName:          arguments.ProjectName,
		ProjectBasePath:      arguments.ProjectBasePath,
		ProjectPath:          arguments.ProjectBasePath,
		ApiFileName:          arguments.ApiFileName,
		ApiInfoFileName:      arguments.ApiInfoFileName,
		DatabaseFileName:     arguments.DatabaseFileName,
		DatabaseInfoFileName: arguments.DatabaseInfoFileName,
		ServiceFileName:      arguments.ServiceFileName,
		OrganizationName:     arguments.OrganizationName,
		OpenAPISpec:          nil,
		DatabaseSchema:       nil,
		TypeMapper:           nil,
		Count:                0,
		Debug:                arguments.Debug,
		HasAdminAPI:          !arguments.NoAdmin,
	}

	_processes := []processes.ProcessInterface{
		&a_download_template.Process{},
		&b_parse_api_spec_file.Process{},
		&b_parse_api_info_file.Process{},
		&b_parse_service_definition_file.Process{},
		&c_parse_database_spec_file.Process{},
		&c_parse_database_info_file.Process{},
		&d_build_crossover_information.Process{},
		&e_build_app_api.Process{},
		&f_build_database.Process{},
	}
	if !arguments.NoAdmin {
		_processes = append(_processes, &g_build_admin_api.Process{})
	}

	for index, process := range _processes {
		fmt.Println("Running process", index+1, "of", len(_processes))
		_payload, err := process.Execute(_payload)
		if err != nil {
			red := color.New(color.FgRed)
			_, _ = red.Println(err.Error())
			panic(err)
		}
		_payload.Count++
	}

	fmt.Println(_payload.ProjectName + ": Project created successfully")

	return nil
}
