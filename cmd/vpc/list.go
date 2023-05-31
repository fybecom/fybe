package cmd

import (
	"context"
	"encoding/json"
	"strconv"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vpcsGetCmd = &cobra.Command{
	Use:     "vpcs",
	Short:   "All about your VPCs.",
	Long:    `Retrieves information about one or multiple VPCs. Filter by name.`,
	Example: `fybe get vpcs`,
	Run: func(cmd *cobra.Command, args []string) {
		ApiRetrieveVpcListRequest := client.ApiClient().
			PrivateNetworksApi.RetrievePrivateNetworkList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if listVpnNameFilter != "" {
			ApiRetrieveVpcListRequest = ApiRetrieveVpcListRequest.Name(listVpnNameFilter)
		}

		resp, httpResp, err := ApiRetrieveVpcListRequest.Execute()

		util.HandleErrors(err, httpResp, "while retrieving VPCs")

		responseJson, _ := json.Marshal(resp.Data)

		if viper.GetString("output") != "json" && viper.GetString("output") != "yaml" {
			var vpcs []*PrivateNetwork
			for _, vpn := range resp.Data {
				vpc := SetVpc(vpn.PrivateNetworkId, vpn.Name, vpn.Description, vpn.RegionName, vpn.DataCenter, vpn.Cidr, vpn.AvailableIps)
				var len = len(vpn.Instances)
				// if instance list has elements
				if len > 0 {
					var instanceIds []string
					for _, instance := range vpn.Instances {
						instanceIds = append(instanceIds, strconv.FormatInt(instance.InstanceId, 10))
					}
					vpc.Instances = instanceIds
				}
				vpcs = append(vpcs, vpc)
			}
			responseJson, _ = json.Marshal(vpcs)
		}

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"vpcId", "name", "description", "regionName", "dataCenter"},
			WideFilter: []string{"vpcId", "name", "description", "regionName", "dataCenter", "cidr", "availableIps", "instances"},
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

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		listVpnNameFilter = viper.GetString("name")

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(vpcsGetCmd)

	vpcsGetCmd.Flags().StringVarP(&listVpnNameFilter, "name", "n", "",
		`Filter by VPC name`)
}
