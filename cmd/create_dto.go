package cmd

import (
	createCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create"
	"github.com/spf13/cobra"
)

var createDtoArguments = createCommand.Arguments{
	Type:              "dto",
	Name:              "",
	RelatedModelNames: []string{},
	RelatedResponse:   nil,
	ProjectPath:       "",
	Debug:             false,
}

var createDtoCmd = &cobra.Command{
	Use:   "dto",
	Short: "Create a new dto",
	Long: `Create a new resource on the project.:

rocket create dto your-dto-name
`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createDtoArguments.Name = name
		command := createCommand.SubCommand{}
		err := command.Execute(createDtoArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	createCmd.AddCommand(createDtoCmd)
	createDtoCmd.Flags().StringVarP(&createDtoArguments.ProjectPath, "path", "p", "", "path to create project")
}
