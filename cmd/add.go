package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new resource",
	Long: `Add new resource to the project. Such as service, package, etc.:

rocket add service your-service-name
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
