package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var AssignInstanceCmd = &cobra.Command{
	Use:   "assign",
	Short: "Add compute instance to VPC",
	Long:  `Add a specific compute instance to a specific VPC using their ips`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
	Args:       cobra.OnlyValidArgs,
	SuggestFor: []string{"assign", "add", "addInstance"},
	ValidArgs:  []string{"vpc"},
}

func init() {
	rootCmd.AddCommand(AssignInstanceCmd)
}
