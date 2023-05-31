package cmd

import (
	"context"
	"fmt"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var generateSecretUserCmd = &cobra.Command{
	Use:     "user",
	Short:   "Generate client secret",
	Long:    `Generate a new client secret and return it`,
	Example: `fybe generateSecret user`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().UsersApi.
			GenerateClientSecret(context.Background()).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while generating client secret")

		clientSecret := resp.Data[0].Secret

		if resp.Data != nil {
			config := cliCmd.ReadConfigFile()
			config.Oauth2ClientSecret = clientSecret
			cliCmd.SaveConfigFile(config)
		}

		fmt.Printf("%v\n", clientSecret)
	},
}

func init() {
	cliCmd.GenerateCmd.AddCommand(generateSecretUserCmd)
}
