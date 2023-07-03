package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var tagAssignmentGetCmd = &cobra.Command{
	Use:   "tagAssignment [tagId] [resourceType] [resourceId]",
	Short: "Info about a assignments of a specif tag",
	Long: `Retrives information about assignments of a specifu identified by
	tag assignment id, resource type and resource id`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().TagAssignmentsApi.
			RetrieveAssignment(context.Background(), tagId, resourceType, resourceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving tag assignment")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"tagId", "tagName", "resourceType", "resourceId", "resourceName"},
			WideFilter: []string{"tagId", "tenantId", "customerId", "tagId", "tagName", "resourceType", "resourceId", "resourceName"},
			JsonPath:   cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 3 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 3 {
			cmd.Help()
			log.Fatal("Too little arguments please specify tagId, resourceType and resourceId")
		}
		tagId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided tagId %v is not valid.", args[0]))
		}

		tagId = tagId64
		resourceType = args[1]
		resourceId = args[2]

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(tagAssignmentGetCmd)
}
