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

var tagsGetCmd = &cobra.Command{
	Use:   "tags",
	Short: "All about your tags",
	Long:  `Retrieves information about one or multiple tags. Filter by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveTagListRequest := client.ApiClient().
			TagsApi.RetrieveTagList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if listTagNameFilter != "" {
			ApiRetrieveTagListRequest = ApiRetrieveTagListRequest.Name(listTagNameFilter)
		}

		resp, httpResp, err := ApiRetrieveTagListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving tags")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"tagId", "name", "color"},
			WideFilter: []string{"tagId", "name", "color"},
			JsonPath:   cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("tagName", cmd.Flags().Lookup("tagName"))
		listTagNameFilter = viper.GetString("tagName")

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(tagsGetCmd)

	tagsGetCmd.Flags().StringVarP(&listTagNameFilter, "tagName", "t", "",
		`Filter by tag name`)
}
