package cmd

import (
	createCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	"github.com/spf13/cobra"
)

var createServiceArguments = createCommand.Arguments{
	Type:              "service",
	Name:              "",
	RelatedModelNames: []string{},
	RelatedResponse:   nil,
	ProjectPath:       "",
	Debug:             false,
}

var createServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a new service",
	Long: `Create a new resource on the project.:

rocket create service your-service-name
`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createServiceArguments.Name = name
		command := createCommand.SubCommand{}
		err := command.Execute(createServiceArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	createCmd.AddCommand(createServiceCmd)
	createServiceCmd.Flags().StringVarP(&createServiceArguments.ProjectPath, "path", "p", "", "path to create project")
}
