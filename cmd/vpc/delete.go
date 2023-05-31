package cmd

import (
	"context"
	"fmt"
	"strconv"

	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var vpcDeleteCmd = &cobra.Command{
	Use:   "vpc [vpcId]",
	Short: "Deletes a specific VPC by id",
	Long:  `Specify a VPC id to delete. All the instances will be unassigned form this VPC`,
	Run: func(cmd *cobra.Command, args []string) {
		httpResp, err := client.ApiClient().
			PrivateNetworksApi.DeletePrivateNetwork(context.Background(), deletevpcId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while deleting VPC")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a vpcId")
		}

		vpcId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided vpcId %v is not valid.", args[0]))
		}
		deletevpcId = vpcId64

		return nil
	},
}

func init() {
	cliCmd.DeleteCmd.AddCommand(vpcDeleteCmd)
}
