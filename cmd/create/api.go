package create

import (
	command "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/api"
	"github.com/spf13/cobra"
)

var createApiArguments = command.Arguments{
	Type:              "api",
	Name:              "",
	RelatedModelNames: []string{},
	RelatedResponse:   nil,
	ProjectPath:       "",
	Debug:             false,
}

var createApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Create a new api",
	Long: `Create a new resource on the project.:

rocket create api your-api-name
`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createApiArguments.Name = name
		command := command.Command{}
		err := command.Execute(createApiArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	Cmd.AddCommand(createApiCmd)
	createApiCmd.Flags().StringVarP(&createApiArguments.ProjectPath, "path", "p", "", "path to create project")
}
