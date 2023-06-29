package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	objectStoragesClient "fybe.com/cli/fybe/apiclient"
	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var objectStorageCreateCmd = &cobra.Command{
	Use:   "objectStorage",
	Short: "Creates a new objectStorage.",
	Long:  `Creates a new objectStorage based on json / yaml input or arguments.`,
	Example: `fybe create objectStorage --region us-central-1 --displayName example --totalPurchasedSpaceTB 1.00 --scalingState enabled --scalingLimitTB 1.00`,
	Run: func(cmd *cobra.Command, args []string) {
		createObjectStorageRequest := *objectStoragesClient.NewCreateObjectStorageRequestWithDefaults()
		content := cliCmd.OpenStdinOrFile()

		switch content {
		case nil:

			// If we dont set createScalingLimit on a disabled scaling state we recieve an error 
			if createScalingState == "disabled" && createScalingLimitTB == nil {
				createScalingLimitTB = 1.00
			}

			autoScaling := objectStoragesClient.AutoScalingTypeRequest{
				State:       createScalingState,
				SizeLimitTB: createScalingLimitTB,
			}

			createObjectStorageRequest.DisplayName = &createdisplayName
			if createdisplayName == "" {
				createObjectStorageRequest.DisplayName = nil
			}

			createObjectStorageRequest.AutoScaling = &autoScaling
			createObjectStorageRequest.Region = createRegion
			createObjectStorageRequest.TotalPurchasedSpaceTB = createTotalPurchasedSpaceTB

		default:
			// from file / stdin
			var requestFromFile objectStoragesClient.CreateObjectStorageRequest
			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge createObjectStorageRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).Decode(&createObjectStorageRequest)
		}

		resp, httpResp, err := client.ApiClient().ObjectStoragesApi.
			CreateObjectStorage(context.Background()).XRequestId(uuid.NewV4().String()).
			CreateObjectStorageRequest(createObjectStorageRequest).Execute()

		util.HandleErrors(err, httpResp, "while creating objectStorage")

		fmt.Printf("%v\n", resp.Data[0].ObjectStorageId)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateCreateInput()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("region", cmd.Flags().Lookup("region"))
		createRegion = viper.GetString("region")

		viper.BindPFlag("displayName", cmd.Flags().Lookup("displayName"))
		createdisplayName = viper.GetString("displayName")

		viper.BindPFlag("scalingState", cmd.Flags().Lookup("scalingState"))
		createScalingState = viper.GetString("scalingState")

		if !(createScalingState == "enabled" || createScalingState == "disabled") {
			cmd.Help()
			log.Fatal("Autoscaling state can only be enabled or disabled.")
		}

		viper.BindPFlag("scalingLimitTB", cmd.Flags().Lookup("scalingLimitTB"))
		createScalingLimitTB = viper.GetFloat64("scalingLimitTB")

		viper.BindPFlag("totalPurchasedSpaceTB", cmd.Flags().Lookup("totalPurchasedSpaceTB"))
		createTotalPurchasedSpaceTB = viper.GetFloat64("totalPurchasedSpaceTB")
		return nil
	},
}

func init() {
	cliCmd.CreateCmd.AddCommand(objectStorageCreateCmd)

	objectStorageCreateCmd.Flags().StringVarP(&createRegion, "region", "r", "", `Region where the objectStorage gets deployed.`)

	objectStorageCreateCmd.Flags().StringVarP(&createdisplayName, "displayName", "n", "", `Display name helps to differentiate between object storages.`)

	objectStorageCreateCmd.Flags().StringVarP(&createScalingState, "scalingState", "s", "disabled",
		`Set scalingState to enable autoscaling (allowed values are enabled|disabled)`)

	objectStorageCreateCmd.Flags().Float64VarP(&createScalingLimitTB, "scalingLimitTB", "l", 0,
		`The size limit for autoscaling in TB, required if scalingState is set.`)

	objectStorageCreateCmd.Flags().Float64VarP(&createTotalPurchasedSpaceTB, "totalPurchasedSpaceTB", "t", 0,
		`The amount of storage pruchased, in TB.`)
}
