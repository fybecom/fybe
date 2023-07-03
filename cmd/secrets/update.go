package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	secretClient "fybe.com/cli/fybe/apiclient"
	"fybe.com/cli/fybe/client"
	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/cmd/util"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var secretUpdateCmd = &cobra.Command{
	Use:   "secret [secretId]",
	Short: "Update a specific secret",
	Long:  `Updates a specific secret based on json / yaml input or arguments.`,
	Example: `fybe update secret 21 --name 'First Secret' ` +
		`--value 'secret'`,
	Run: func(cmd *cobra.Command, args []string) {
		updateSecretRequest := *secretClient.NewUpdateSecretRequestWithDefaults()
		content := cliCmd.OpenStdinOrFile()

		switch content {
		case nil:
			if updateSecretName != "" {
				updateSecretRequest.Name = &updateSecretName
			}
			if updateSecretValue != "" {
				updateSecretRequest.Value = &updateSecretValue
			}
		default:
			// from file / stdin
			var requestFromFile secretClient.UpdateSecretRequest

			err := json.Unmarshal(content, &requestFromFile)
			if err != nil {
				log.Fatal(fmt.Sprintf("Format invalid. Please check your syntax: %v", err))
			}
			// merge updateSecretRequest with one from file to have the defaults there
			json.NewDecoder(strings.NewReader(string(content))).
				Decode(&updateSecretRequest)
		}

		resp, httpResp, err := client.ApiClient().SecretsApi.
			UpdateSecret(context.Background(), updateSecretId).
			UpdateSecretRequest(updateSecretRequest).
			XRequestId(uuid.NewV4().String()).
			Execute()

		util.HandleErrors(err, httpResp, "while updating secret")

		responseJSON, _ := resp.MarshalJSON()
		log.Info(fmt.Sprintf("%v", string(responseJSON)))
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateCreateInput()

		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}
		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a secretId.")
		}

		viper.BindPFlag("name", cmd.Flags().Lookup("name"))
		updateSecretName = viper.GetString("name")

		viper.BindPFlag("value", cmd.Flags().Lookup("value"))
		updateSecretValue = viper.GetString("value")

		secretId64, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided secretId %v is not valid.", args[0]))
		}
		updateSecretId = secretId64

		return nil
	},
}

func init() {
	cliCmd.UpdateCmd.AddCommand(secretUpdateCmd)

	secretUpdateCmd.Flags().StringVar(&updateSecretName, "name", "",
		`name of the secret`)

	secretUpdateCmd.Flags().StringVar(&updateSecretValue, "value", "",
		`value of the secret`)
}
