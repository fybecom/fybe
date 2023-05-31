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

type jmap map[string]interface{}

var instanceGetCmd = &cobra.Command{
	Use:   "instance [instanceId]",
	Short: "Info about a specific instance.",
	Long:  `Retrieves information about an instance identified by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().InstancesApi.
			RetrieveInstance(context.Background(), getInstanceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving instance")

		arr := make([]jmap, 0)
		for _, entry := range resp.Data {
			entryModified, _ := util.StructToMap(entry)
			if &entry.IpConfig != nil {
				entryModified["ipv4"] = entry.IpConfig.V4.Ip
				entryModified["ipv6"] = entry.IpConfig.V6.Ip
			} else {
				entryModified["ipv4"] = ""
				entryModified["ipv6"] = ""
			}
			arr = append(arr, entryModified)
		}

		responseJson, _ := json.Marshal(arr)

		configFormatter := outputFormatter.FormatterConfig{
			Filter: []string{"instanceId", "name", "displayName", "status", "imageId", "ipv4", "ipv6", "defaultUser"},
			WideFilter: []string{
				"instanceId", "name", "displayName", "status", "imageId", "region", "productId", "customerId", "ipv4", "ipv6", "defaultUser",
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
			log.Fatal("Please provide an instanceId.")
		}
		instanceId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided instanceId %v is not valid.", args[0]))
		}
		getInstanceId = instanceId64

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(instanceGetCmd)
}
