package cmd

import (
	"context"
	"encoding/json"
	"os"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/outputFormatter"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var regionGetCmd = &cobra.Command{
	Use:   "datacenters",
	Short: "List all data centers",
	Long:  `Retrieves all available data centers.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveDatacenterListRequest := client.ApiClient().
			ObjectStoragesApi.RetrieveDataCenterList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})
		resp, httpResp, err := ApiRetrieveDatacenterListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving datacenters")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{

			Filter:     []string{"slug", "capabilities", "regionName", "regionSlug"},
			WideFilter: []string{"name", "slug", "capabilities", "s3Url", "regionName", "regionSlug"},
			JsonPath:   cliCmd.OutputFormatDetails}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()
		if len(args) > 1 {
			cmd.Help()
			os.Exit(0)
		}

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(regionGetCmd)
}
