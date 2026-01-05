package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRolePermissionsCmd = &cobra.Command{
	Use:     "list-role-permissions [role_id]",
	Aliases: []string{"list-permissions"},
	Short:   "List permissions for a role",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PermissionsResponse
		var err error

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.ListRolePermissions(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-role-permissions")

		cmd.Println(cmdutil.FormatJSON(res, "permissions"))
	},
}

func init() {
	Cmd.AddCommand(ListRolePermissionsCmd)
}
