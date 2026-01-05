package users

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListUserPermissionsCmd = &cobra.Command{
	Use:     "list-user-permissions [user_id]",
	Aliases: []string{"list-permissions"},
	Short:   "Get a user permissions",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.ListUserPermissions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-user-permissions")

		cmd.Println(cmdutil.FormatJSON(res, "permissions"))
	},
}

func init() {
	Cmd.AddCommand(ListUserPermissionsCmd)
}
