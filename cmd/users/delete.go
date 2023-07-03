package cmd

import (
	"context"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userDeleteCmd = &cobra.Command{
	Use:   "user [userId]",
	Short: "Deletes a specific user",
	Long:  `Specify a tag id to delete the specified user.`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().UsersApi.
			DeleteUser(context.Background(), deleteUserId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting user")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an userId.")
		}

		deleteUserId = args[0]

		return nil
	},
}

func init() {
	cliCmd.DeleteCmd.AddCommand(userDeleteCmd)
}
