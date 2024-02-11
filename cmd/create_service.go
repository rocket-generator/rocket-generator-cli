package cmd

import (
	"github.com/rocket-generator/rocket-generator-cli/modules/commands/create/service"
	createServiceCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/service"
	"github.com/spf13/cobra"
)

var createServiceArguments = service.Arguments{
	ServiceName:       "",
	RelatedModelNames: []string{},
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
		serviceName := args[0]
		createServiceArguments.ServiceName = serviceName
		command := createServiceCommand.Command{}
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
