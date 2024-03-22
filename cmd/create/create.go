package create

import (
	"github.com/rocket-generator/rocket-generator-cli/cmd"
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
	cmd.RootCmd.AddCommand(Cmd)
}
