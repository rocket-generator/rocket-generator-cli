package create

import (
	command "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/dto"
	"github.com/spf13/cobra"
)

var createDtoArguments = command.Arguments{
	Type:              "dto",
	Name:              "",
	RelatedModelNames: []string{},
	RelatedResponse:   nil,
	ProjectPath:       "",
	Debug:             false,
}

var DtoCmd = &cobra.Command{
	Use:   "dto",
	Short: "Create a new dto",
	Long: `Create a new resource on the project.:

rocket create dto your-dto-name
`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createDtoArguments.Name = name
		command := command.Command{}
		err := command.Execute(createDtoArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	DtoCmd.Flags().StringVarP(&createDtoArguments.ProjectPath, "path", "p", "", "path to create project")
}
