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

var objectStorageHistoryCmd = &cobra.Command{
	Use:     "objectStorages",
	Short:   "History of your object storage",
	Long:    `Show what happend to your object storage over time.`,
	Example: `fybe history objectStorages --objectStorageId 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().ObjectStoragesAuditsApi.
			RetrieveObjectStorageAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if historyObjectStorageIdFilter != "" {
			historyRequest = historyRequest.ObjectStorageId(historyObjectStorageIdFilter)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving object storage history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"id", "objectStorageId", "action", "username", "timestamp"},
			WideFilter: []string{"id", "objectStorageId", "action", "username", "changedBy", "requestId", "traceId", "timestamp", "changes"},
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

		viper.BindPFlag("objectStorageId", cmd.Flags().Lookup("objectStorageId"))
		historyObjectStorageIdFilter = viper.GetString("objectStorageId")

		return nil
	},
}

func init() {
	cliCmd.HistoryCmd.AddCommand(objectStorageHistoryCmd)

	objectStorageHistoryCmd.Flags().StringVarP(&historyObjectStorageIdFilter, "objectStorageId", "", "",
		`to filter audits using object storage id`)
}
