package cmd

import (
	"context"
	"strconv"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var removeInstanceFromvpcCmd = &cobra.Command{
	Use:     "vpc [vpcId] [instanceId]",
	Short:   "Remove instance from VPC",
	Long:    `Remove a specific instance from a specific VPC using their ips`,
	Example: `fybe unassign vpc 12345 100`,
	Run: func(cmd *cobra.Command, args []string) {
		_, httpResp, err := client.ApiClient().PrivateNetworksApi.
			UnassignInstancePrivateNetwork(context.Background(), unassignvpcId, unassignInstanceId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while removing from VPC")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateCreateInput()

		if len(args) > 2 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 2 {
			cmd.Help()
			log.Fatal("Please provide an vpcId and instanceId.")
		}

		vpcIdInt64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		assignInstanceIdInt64, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		unassignvpcId = vpcIdInt64
		unassignInstanceId = assignInstanceIdInt64

		return nil
	},
}

func init() {
	cliCmd.UnassignInstanceCmd.AddCommand(removeInstanceFromvpcCmd)
}
