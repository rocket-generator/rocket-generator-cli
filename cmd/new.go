package cmd

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var arguments = newCommand.Arguments{
	ProjectName:      "",
	ProjectBasePath:  "",
	Template:         "",
	ApiFileName:      "",
	DatabaseFileName: "",
	OrganizationName: "",
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Long: `Create new project with specified template.

rocket new your-service-name --template go-gin
`,
	Run: func(cmd *cobra.Command, args []string) {
		arguments.ProjectName = args[0]
		if arguments.ProjectBasePath == "" {
			currentPath, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
				return
			}
			arguments.ProjectBasePath = currentPath
		}

		command := newCommand.Command{}
		err := command.Execute(arguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&arguments.Template, "template", "t", "go-gin", "specify template to use")
	newCmd.Flags().StringVarP(&arguments.ApiFileName, "api", "a", "api.yaml", "specify OpenAPI Spec Yaml file")
	newCmd.Flags().StringVarP(&arguments.DatabaseFileName, "database", "d", "api.yaml", "specify database PlantUML file")
	newCmd.Flags().StringVarP(&arguments.OrganizationName, "organization", "o", "your_org", "specify your (github) organization name")
	newCmd.Flags().StringVarP(&arguments.ProjectBasePath, "path", "p", "", "path to create project")

}
