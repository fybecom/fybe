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
)

var objectStorageStatsCmd = &cobra.Command{
	Use:     "objectStorage [objectStorageId]",
	Short:   "get statistics of objectStorage.",
	Long:    `get statistics of objectStorage based on its id.`,
	Example: `fybe stats objectStorage 1f771979-1c0f-44ab-ab5b-2c3752731b45`,
	Run: func(cmd *cobra.Command, args []string) {

		resp, httpResp, err := client.ApiClient().ObjectStoragesApi.
			RetrieveObjectStoragesStats(context.Background(), statsObjectStorageId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving object storage")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"usedSpaceTB", "usedSpacePercentage", "numberOfObjects"},
			WideFilter: []string{"tenantId", "customerId", "usedSpaceTB", "usedSpacePercentage", "numberOfObjects"},
			JsonPath:   cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an objectStorageId.")
		}

		statsObjectStorageId = args[0]
		return nil
	},
}

func init() {
	cliCmd.StatsCmd.AddCommand(objectStorageStatsCmd)
}
