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

var secretsHistoryCmd = &cobra.Command{
	Use:   "secrets",
	Short: "History of your secrets",
	Long:  `Show what happend to your secrets over time.`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().
			SecretsAuditsApi.RetrieveSecretAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if historySecretIdFilter != 0 {
			historyRequest = historyRequest.SecretId(historySecretIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving secrets history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "secretId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "secretId", "action", "username", "timestamp", "tenantId", "customerId", "changedBy", "requestId", "traceId", "changes"},
			JsonPath:   cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("secretId", cmd.Flags().Lookup("secretId"))
		historySecretIdFilter = viper.GetInt64("secretId")

		return nil
	},
}

func init() {
	cliCmd.HistoryCmd.AddCommand(secretsHistoryCmd)

	secretsHistoryCmd.Flags().Int64Var(&historySecretIdFilter, "secretId", 0,
		`To filter audits using Secret Id`)
}
