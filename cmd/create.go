package cmd

import (
	"github.com/rocket-generator/rocket-generator-cli/cmd/create"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create new resource",
	Long: `Create new resource on the project. Such as service, package, etc.:

rocket create service your-service-name
`,
}

func init() {
	RootCmd.AddCommand(Cmd)
	Cmd.AddCommand(create.ApiCmd)
	Cmd.AddCommand(create.ModelCmd)
	Cmd.AddCommand(create.DtoCmd)
	Cmd.AddCommand(create.ServiceCmd)
}
