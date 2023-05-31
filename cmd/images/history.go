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
	Use:     "images",
	Short:   "History of your images",
	Long:    `Show what happend to your images over time.`,
	Example: `fybe history images --imageId 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d`,
	Run: func(cmd *cobra.Command, args []string) {
		historyRequest := client.ApiClient().ImagesAuditsApi.
			RetrieveImageAuditsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if historyImageIdFilter != "" {
			historyRequest = historyRequest.ImageId(
				historyImageIdFilter,
			)
		}

		resp, httpResp, err := historyRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving image history")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"id", "imageId", "action", "timestamp",
			},
			WideFilter: []string{
				"id", "imageId", "action", "username", "changedBy",
				"requestId", "traceId", "timestamp", "changes",
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

		viper.BindPFlag("imageId", cmd.Flags().Lookup("imageId"))
		historyImageIdFilter = viper.GetString("imageId")

		return nil
	},
}

func init() {
	cliCmd.HistoryCmd.AddCommand(historyCmd)

	historyCmd.Flags().StringVarP(&historyImageIdFilter, "imageId", "i", "",
		`Filter by a specific image via its imageId.`)
}
