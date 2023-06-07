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

// imageGetCmd represents the get command
var imageGetCmd = &cobra.Command{
	Use:     "image [imageId]",
	Short:   "Get data about an image",
	Long:    `Retrieves information about one image by its id.`,
	Example: `fybe get image 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().
			ImagesApi.RetrieveImage(context.Background(), getImageId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving image")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{
				"imageId", "name", "sizeMb", "uploadedSizeMb", "osType", "version", "status",
			},
			WideFilter: []string{
				"imageId", "name", "sizeMb", "uploadedSizeMb", "osType", "version",
				"status", "standardImage", "url", "description", "errorMessage",
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

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide an imageId.")
		}

		getImageId = args[0]

		if getImageId == "" {
			cmd.Help()
			log.Fatal("Argument imageId is empty. Please provide a non empty imageId.")
		}

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(imageGetCmd)
}
