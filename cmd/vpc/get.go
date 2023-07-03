package cmd

import (
	"context"
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

var vpcGetCmd = &cobra.Command{
	Use:     "vpc [vpcId]",
	Short:   "Info about a specific vpc",
	Long:    `Retrieves information about one vpc identified by id.`,
	Example: `fybe get vpc 12`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().VirtualPrivateCloudVPCApi.
			RetrievePrivateNetwork(context.Background(), getvpcId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving VPC")

		responseJson, _ := mapInstancesToIds(resp.Data)

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
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a vpcId.")
		}

		vpcId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided vpcId %v is not valid.", args[0]))
		}
		getvpcId = vpcId64

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(vpcGetCmd)
}
