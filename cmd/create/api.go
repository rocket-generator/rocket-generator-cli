package create

import (
	command "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/api"
	"github.com/spf13/cobra"
	"github.com/stoewer/go-strcase"
)

var createApiArguments = command.Arguments{
	Type:            "api",
	Path:            "",
	Method:          "",
	ApiFileName:     "",
	ApiInfoFileName: "",
	ProjectPath:     "",
	ServiceFileName: "",
	Debug:           false,
}

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Create a new api",
	Long: `Create a new resource on the project.:

rocket create api get /users
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			panic("Please provide the name of the api")
		}
		method := args[0]
		path := args[1]
		createApiArguments.Method = strcase.SnakeCase(method)
		createApiArguments.Path = path
		command := command.Command{}
		err := command.Execute(createApiArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	ApiCmd.Flags().StringVarP(&createApiArguments.ProjectPath, "path", "p", "", "path to create project")
	ApiCmd.Flags().StringVarP(&createApiArguments.ApiFileName, "api", "a", "api.yaml", "specify OpenAPI Spec Yaml file")
	ApiCmd.Flags().StringVarP(&createApiArguments.DatabaseFileName, "database", "d", "api.yaml", "specify database PlantUML file")
	ApiCmd.Flags().StringVarP(&createApiArguments.ApiInfoFileName, "api_info", "i", "", "specify api info json file")
	ApiCmd.Flags().StringVarP(&createApiArguments.DatabaseInfoFileName, "db_info", "j", "", "specify database info json file")
	ApiCmd.Flags().StringVarP(&createApiArguments.ServiceFileName, "service", "s", "", "specify service path mapping json file")
}
