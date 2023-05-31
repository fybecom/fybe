package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var secretGetCmd = &cobra.Command{
	Use:     "secret [secretId]",
	Short:   "Info about a specific secret",
	Long:    `Get data about a specific secret on your account`,
	Example: `fybe get secret 21`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSecretRequest := client.ApiClient().
			SecretsApi.RetrieveSecret(context.Background(), getSecretId).
			XRequestId(uuid.NewV4().String())

		resp, httpResp, err := ApiRetrieveSecretRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving secret")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"secretId", "type", "name",
			},
			WideFilter: []string{
				"secretId", "type", "name", "customerId", "tenantId",
			},
			JsonPath: cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a secretId.")
		}

		secretId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Specified secretId %v is not valid.", args[0]))
		}
		getSecretId = secretId64

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(secretGetCmd)
}
