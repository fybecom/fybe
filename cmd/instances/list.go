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

var instancesGetCmd = &cobra.Command{
	Use:     "instances",
	Short:   "List your instances",
	Long:    `Retrieves information about one or multiple instances. Filter by name.`,
	Example: `fybe get instances`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveInstanceListRequest := client.ApiClient().
			InstancesApi.RetrieveInstancesList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if listInstanceNameFilter != "" {
			ApiRetrieveInstanceListRequest = ApiRetrieveInstanceListRequest.Name(listInstanceNameFilter)
		}

		resp, httpResp, err := ApiRetrieveInstanceListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving instances")

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
			Filter: []string{
				"instanceId", "name", "displayName", "status", "imageId", "region", "productId", "ipv4", "ipv6", "defaultUser",
			},
			WideFilter: []string{
				"instanceId", "name", "displayName", "status", "imageId", "region", "productId", "customerId", "tenantId", "ipv4", "ipv6", "defaultUser",
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

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listInstanceNameFilter = viper.GetString("name")

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(instancesGetCmd)

	instancesGetCmd.Flags().StringVarP(&listInstanceNameFilter, "name", "n", "",
		`Filter by instance name`)
}
