package cmd

import (
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var newArguments = newCommand.Arguments{
	ProjectName:      "",
	ProjectBasePath:  "",
	Template:         "",
	ApiFileName:      "",
	ApiInfoFileName:  "",
	DatabaseFileName: "",
	ServiceFileName:  "",
	OrganizationName: "",
	Debug:            false,
	NoAdmin:          false,
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new project",
	Long: `Create new project with specified template.

rocket new your-service-name --template go-gin
`,
	Run: func(cmd *cobra.Command, args []string) {
		newArguments.ProjectName = args[0]
		if newArguments.ProjectBasePath == "" {
			currentPath, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
				return
			}
			newArguments.ProjectBasePath = currentPath
		}
		newArguments.Debug = debugFlag

		command := newCommand.Command{}
		err := command.Execute(newArguments)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&newArguments.Template, "template", "t", "go-gin", "specify template to use")
	newCmd.Flags().StringVarP(&newArguments.ApiFileName, "api", "a", "api.yaml", "specify OpenAPI Spec Yaml file")
	newCmd.Flags().StringVarP(&newArguments.DatabaseFileName, "database", "d", "api.yaml", "specify database PlantUML file")
	newCmd.Flags().StringVarP(&newArguments.ApiInfoFileName, "api_info", "i", "", "specify api info json file")
	newCmd.Flags().StringVarP(&newArguments.ServiceFileName, "service", "s", "", "specify service path mapping json file")
	newCmd.Flags().StringVarP(&newArguments.OrganizationName, "organization", "o", "your_org", "specify your (github) organization name")
	newCmd.Flags().StringVarP(&newArguments.ProjectBasePath, "path", "p", "", "path to create project")
	newCmd.Flags().BoolVarP(&newArguments.NoAdmin, "noadmin", "n", false, "no need admin")
}
