package create

import (
	command "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/response"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
)

var createResponseArguments = command.Arguments{
	Type:        "response",
	Name:        "",
	ApiFileName: "",
	ProjectPath: "",
	Debug:       false,
}

var ResponseCmd = &cobra.Command{
	Use:   "api",
	Short: "Create a new response",
	Long: `Create a new resource on the project.:

rocket create api get /users
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			panic("Please provide the name of the api")
		}
		name := args[0]
		createResponseArguments.Name = strcase.SnakeCase(name)
		command := command.Command{}
		err := command.Execute(createResponseArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	ResponseCmd.Flags().StringVarP(&createApiArguments.ProjectPath, "path", "p", "", "path to create project")
	ResponseCmd.Flags().StringVarP(&createApiArguments.ApiFileName, "api", "a", "api.yaml", "specify OpenAPI Spec Yaml file")
	ResponseCmd.Flags().StringVarP(&createApiArguments.DatabaseFileName, "database", "d", "api.yaml", "specify database PlantUML file")
	ResponseCmd.Flags().StringVarP(&createApiArguments.ApiInfoFileName, "api_info", "i", "", "specify api info json file")
	ResponseCmd.Flags().StringVarP(&createApiArguments.DatabaseInfoFileName, "db_info", "j", "", "specify database info json file")
	ResponseCmd.Flags().StringVarP(&createApiArguments.ServiceFileName, "service", "s", "", "specify service path mapping json file")
}
