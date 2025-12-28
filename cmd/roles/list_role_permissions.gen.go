package roles

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRolePermissionsCmd = &cobra.Command{
	Use:   "list-role-permissions [role_id]",
	Aliases: []string{ "list-permissions", },
	Short: "List permissions for a role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.ListRolePermissions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-role-permissions")

		cmdutil.PrintJSON(res, "permissions")
	},
}

func init() {
	Cmd.AddCommand(ListRolePermissionsCmd)
}
