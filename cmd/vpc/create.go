package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	vpcs "fybe.com/cli/fybe/apiclient"
	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vpcCreateCmd = &cobra.Command{
	Use:     "vpc",
	Short:   "Creates a new VPC",
	Long:    `Creates a new VPC based on json / yaml input or arguments.`,
	Example: `fybe create vpc --name myVPC --region us-east-1 --description "My First VPC"`,
	Run: func(cmd *cobra.Command, args []string) {
		createvpcRequest := *vpcs.NewCreatePrivateNetworkRequestWithDefaults()
		content := cliCmd.OpenStdinOrFile()

		switch content {
		case nil:

			// from arguments
			createvpcRequest.Name = createVpnName
			createvpcRequest.Region = &createVpnRegion
			if (createVpnDescription) != "" {
				createvpcRequest.Description = &createVpnDescription
			}

		default:
			// from file / stdin
			var requestFromFile vpcs.CreatePrivateNetworkRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createvpcRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createvpcRequest)
		}

		resp, httpResp, err := client.ApiClient().VirtualPrivateCloudVPCApi.
			CreatePrivateNetwork(context.Background()).XRequestId(uuid.NewV4().String()).
			CreatePrivateNetworkRequest(createvpcRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating VPC")

		fmt.Printf("%v\n", resp.Data[0].PrivateNetworkId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		createVpnName = viper.GetString("name")

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		createVpnRegion = viper.GetString("region")

		viper.BindPFlag("description", cmd.Flags().Lookup("description"))
		createVpnDescription = viper.GetString("description")

		if cliCmd.InputFile == "" {
			// arguments required
			if createVpnName == "" {
				cmd.Help()
				log.Fatal("Argument name is empty. Please provide one.")
			}
			if createVpnRegion == "" {
				cmd.Help()
				log.Fatal("Argument region is empty. Please provide one.")
			}
		}

		return nil
	},
}

func init() {
	cliCmd.CreateCmd.AddCommand(vpcCreateCmd)

	vpcCreateCmd.Flags().StringVarP(&createVpnName, "name", "n", "", `Name of the VPC.`)
	vpcCreateCmd.Flags().StringVar(&createVpnRegion, "region", "", `Region of the VPC.`)
	vpcCreateCmd.Flags().StringVarP(&createVpnDescription, "description", "", "", `Description of the VPC.`)
}
