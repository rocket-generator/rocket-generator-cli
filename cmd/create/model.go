package create

import (
	command "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/model"
	"github.com/spf13/cobra"
)

var createModelArguments = command.Arguments{
	Type:             "model",
	Name:             "",
	DatabaseFileName: "",
	ProjectPath:      "",
	Debug:            false,
}

var ModelCmd = &cobra.Command{
	Use:   "model",
	Short: "Create a new model and related files",
	Long: `Create a new model on the project.:

rocket create model your-model-name
`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createModelArguments.Name = name
		command := command.Command{}
		err := command.Execute(createModelArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	ModelCmd.Flags().StringVarP(&createModelArguments.ProjectPath, "path", "p", "", "path to create project")
	ModelCmd.Flags().StringVarP(&createModelArguments.DatabaseFileName, "database", "d", "api.yaml", "specify database PlantUML file")
}
