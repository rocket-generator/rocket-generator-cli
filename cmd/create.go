package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new resource",
	Long: `Create new resource on the project. Such as service, package, etc.:

rocket create service your-service-name
`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
