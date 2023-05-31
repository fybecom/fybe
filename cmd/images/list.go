package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var imagesGetCmd = &cobra.Command{
	Use:     "images",
	Short:   "All about your images",
	Long:    `Retrieves information about one or multiple images. Filter by name.`,
	Example: `fybe get images`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveImageListRequest := client.ApiClient().
			ImagesApi.RetrieveImageList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if cmd.Flags().Changed("name") {
			ApiRetrieveImageListRequest = ApiRetrieveImageListRequest.Name(listImageNameFilter)
		}

		imageType, _ := cmd.Flags().GetString("imageType")
		if imageType != "" {
			if imageType == "standard" {
				ApiRetrieveImageListRequest = ApiRetrieveImageListRequest.StandardImage(true)
			} else {
				ApiRetrieveImageListRequest = ApiRetrieveImageListRequest.StandardImage(false)
			}
		}

		resp, httpResp, err := ApiRetrieveImageListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving images")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"imageId", "name", "sizeMb", "uploadedSizeMb", "osType", "version", "status",
			},
			WideFilter: []string{
				"imageId", "name", "sizeMb", "uploadedSizeMb", "osType", "tags", "version",
				"url", "description", "status", "errorMessage", "standardImage",
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

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listImageNameFilter = viper.GetString("name")

		viper.BindPFlag("imageType", cmd.Flags().Lookup("imageType"))
		listImageTypeFilter = viper.GetString("imageType")

		if listImageTypeFilter != "" && listImageTypeFilter != "standard" && listImageTypeFilter != "custom" {
			cmd.Help()
			log.Fatal(
				fmt.Sprintf(
					"Invalid value `%v` for imageType. Please provide one of `standard` or `custom`.",
					listImageTypeFilter,
				))
		}

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(imagesGetCmd)

	imagesGetCmd.Flags().StringVarP(&listImageNameFilter, "name", "n", "",
		`Filter by image name.`)

	imagesGetCmd.Flags().StringVar(&listImageTypeFilter, "imageType", "",
		`Filter by type of image. Available values are 'standard' or 'custom'.`)
}
