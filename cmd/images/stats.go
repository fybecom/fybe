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

// imageStatsGetCmd represents the get command
var imageStatsGetCmd = &cobra.Command{
	Use:     "images-stats",
	Short:   "Info about custom images",
	Long:    `Retrieves information about available and used space for custom images.`,
	Example: `fybe get images-stats`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			ImagesApi.RetrieveCustomImagesStats(context.Background()).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving images stats")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"totalSizeMb", "freeSizeMb", "usedSizeMb", "currentImagesCount",
			},
			WideFilter: []string{
				"totalSizeMb", "freeSizeMb", "usedSizeMb", "currentImagesCount",
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

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(imageStatsGetCmd)
}
