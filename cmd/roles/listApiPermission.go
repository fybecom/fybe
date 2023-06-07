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

var permissionGetCommand = &cobra.Command{
	Use:   "permissions",
	Short: "List services and permissions",
	Long:  `Retrieves information about one or multiple apis and the actions that are allowed to be performed on them. Filter by apiName.`,
	Run: func(cmd *cobra.Command, args []string) {

		ApiPermissionRetrieveList := client.ApiClient().RolesApi.
			RetrieveApiPermissionsList(context.Background()).
			XRequestId(uuid.NewV4().String()).
			Page(cliCmd.Page).
			Size(cliCmd.Size).
			OrderBy([]string{cliCmd.OrderBy})

		if listapipermissionsApiNameFilter != "" {
			ApiPermissionRetrieveList = ApiPermissionRetrieveList.ApiName(listapipermissionsApiNameFilter)
		}

		resp, httpResp, err := ApiPermissionRetrieveList.Execute()

		util.HandleErrors(err, httpResp, "while retrieving api permissions")

		responseJson, _ := json.Marshal(resp.Data)

		configFormatter := outputFormatter.FormatterConfig{
			Filter:     []string{"apiName", "actions"},
			WideFilter: []string{"tenantId", "customerId", "apiName", "actions"},
			JsonPath:   cliCmd.OutputFormatDetails,
		}

		util.HandleResponse(responseJson, configFormatter)
	},
	Args: func(cmd *cobra.Command, args []string) error {
		cliCmd.ValidateOutputFormat()

		if len(args) > 0 {
			cmd.Help()
			log.Fatal("Too many positional arguments.")
		}

		viper.BindPFlag("apiName", cmd.Flags().Lookup("apiName"))
		listapipermissionsApiNameFilter = viper.GetString("apiName")

		return nil
	},
}

func init() {
	cliCmd.GetCmd.AddCommand(permissionGetCommand)

	permissionGetCommand.Flags().StringVarP(&listapipermissionsApiNameFilter, "apiName", "a", "",
		`filter by api name if the type of role is an api role`)
}
