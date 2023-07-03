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

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "instancesActions",
	Short: "History of your instance actions",
	Long:  `Show what actions you took on your instance`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().InstanceActionsAuditsApi.
			RetrieveInstancesActionsAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if historyInstanceId != 0 {
			historyRequest = historyRequest.InstanceId(historyInstanceId)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving instance action history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"id", "instanceId", "action", "username", "timestamp",
			},
			WideFilter: []string{
				"id", "instanceId", "action", "username", "changedBy", "requestId",
				"traceId", "tenantId", "customerId", "timestamp", "changes",
			},
			JsonPath: cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("instanceId", cmd.Flags().Lookup("instanceId"))
		historyInstanceId = viper.GetInt64("instanceId")

		return nil
	},
}

func init() {
	cliCmd.HistoryCmd.AddCommand(historyCmd)

	historyCmd.Flags().Int64VarP(&historyInstanceId, "instanceId", "i", 0,
		`To filter audits using Instance Id`)
}
