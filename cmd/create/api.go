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
}
