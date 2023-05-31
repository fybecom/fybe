package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	vpcClient "fybe.com/cli/fybe/apiclient"
	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	"fybe.com/cli/fybe/outputFormatter"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var userUpdateCmd = &cobra.Command{
	Use:   "vpc [vpcId]",
	Short: "Updates a specific vpc",
	Long:  `Updates the specific vpc by setting new values either by file input or flags / environment variables`,
	Run: func(cmd *cobra.Command, args []string) {
		updateVpcRequest := *vpcClient.NewPatchPrivateNetworkRequestWithDefaults()
		content := cliCmd.OpenStdinOrFile()
		switch content {
		case nil:
			// from arguments
			if updateVpnName != "" {
				updateVpcRequest.Name = &updateVpnName
			}
			if cmd.Flags().Changed("description") {
				updateVpcRequest.Description = &updateVpnDescription
			}
		default:
			// from file / stdin
			var requestFromFile vpcClient.PatchPrivateNetworkRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updatevpcRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&updateVpcRequest)
		}

		resp, httpResp, err := client.ApiClient().PrivateNetworksApi.
			PatchPrivateNetwork(context.Background(), updateVpnId).
			PatchPrivateNetworkRequest(updateVpcRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating vpc")

		responseJson, _ := mapInstancesToIds(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"vpcId", "name", "description", "region", "dataCenter"},
			WideFilter: []string{"vpcId", "name", "description", "region", "dataCenter", "cidr", "availableIps", "instances"},
			JsonPath:   cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateCreateInput()
		cliCmd.ValidateOutputFormat()

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a vpcId.")
		}
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateVpnName = viper.GetString("name")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		updateVpnDescription = viper.GetString("description")

		vpcId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided vpcId %v is not valid.", args[0]))
		}

		updateVpnId = vpcId64

		return nil
	},
}

func init() {
	cliCmd.UpdateCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringVarP(&updateVpnName, "name", "n", "",
		`Name of the VPC.`)

	userUpdateCmd.Flags().StringVar(&updateVpnDescription, "description", "",
		`Description of the VPC.`)
}
