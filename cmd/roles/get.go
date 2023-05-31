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

var roleGetCmd = &cobra.Command{
	Use:   "role [roleId]",
	Short: "Info about a specific role",
	Long:  `Retrieves information about one role identified by id.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, httpResp, err := client.ApiClient().RolesApi.
			RetrieveRole(context.Background(), getRoleId).
			XRequestId(uuid.NewV4().String()).Execute()

		util.HandleErrors(err, httpResp, "while retrieving role")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"roleId", "name"},
			WideFilter: []string{"roleId", "name", "admin", "accessAllResources", "permissions"},
			JsonPath:   cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()

		if len(args) < 1 {
			cmd.Help()
			log.Fatal("Please provide a roleId.")
		}
		if len(args) > 1 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		roleId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(fmt.Sprintf("Provided roleId %v is not valid.", args[0]))
		}
		getRoleId = roleId

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(roleGetCmd)
}
