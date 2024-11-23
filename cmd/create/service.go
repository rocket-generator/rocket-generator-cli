package create

import (
	command "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/service"
	"github.com/spf13/cobra"
)

var createServiceArguments = command.Arguments{
	Type:              "service",
	Name:              "",
	RelatedModelNames: []string{},
	RelatedResponse:   nil,
	ProjectPath:       "",
	IsAuthService:     false,
	Debug:             false,
}

var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a new service",
	Long: `Create a new resource on the project.:

rocket create service your-service-name --model=model1 --model=model2
`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createServiceArguments.Name = name
		command := command.Command{}
		err := command.Execute(createServiceArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	ServiceCmd.Flags().StringVarP(&createServiceArguments.ProjectPath, "path", "p", "", "path to create project")
	ServiceCmd.Flags().StringArrayVar(&createServiceArguments.RelatedModelNames, "model", []string{}, "related model names (can be specified multiple times)")
}
