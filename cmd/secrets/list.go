package cmd

import (
	"context"
	"encoding/json"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var secretsGetCmd = &cobra.Command{
	Use:     "secrets",
	Short:   "All about your secrets",
	Long:    `List and filter all secrets in your account filtered by type or name`,
	Example: `fybe get secrets`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveSecretsListRequest := client.ApiClient().
			SecretsApi.RetrieveSecretList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if listSecretNameFilter != "" {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.
				Name(listSecretNameFilter)
		}

		if listSecretTypeFilter != "" {
			ApiRetrieveSecretsListRequest = ApiRetrieveSecretsListRequest.
				Type_(listSecretTypeFilter)
		}

		resp, httpResp, err := ApiRetrieveSecretsListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving secrets")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"secretId", "name", "type",
			},
			WideFilter: []string{
				"secretId", "name", "type", "customerId", "tenantId",
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

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listSecretNameFilter = viper.GetString("name")

		viper.BindPFlag("type", cmd.Flags().Lookup("type"))
		listSecretTypeFilter = viper.GetString("type")

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(secretsGetCmd)

	secretsGetCmd.Flags().StringVarP(&listSecretNameFilter, "name", "n", "",
		`Filter by secret name`)

	secretsGetCmd.Flags().StringVarP(&listSecretTypeFilter, "type", "t", "",
		`Filter by secret type [ssh|password]`)
}
