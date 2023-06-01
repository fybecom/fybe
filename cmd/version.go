package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "snapshot"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version and exits",
	Long: `Shows the current version of fybe CLI.

For updates please refer to http://fybe.com/en/cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("fybe %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
