package users

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListUsersCmd = &cobra.Command{
	Use:     "list-users",
	Aliases: []string{"list"},
	Short:   "List all users",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.ListUsers(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-users")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {
	Cmd.AddCommand(ListUsersCmd)
}
