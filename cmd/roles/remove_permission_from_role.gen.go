package roles

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RemovePermissionFromRoleCmd = &cobra.Command{
	Use:     "remove-permission-from-role [role_id]",
	Aliases: []string{"remove-permission-from"},
	Short:   "Revoke permission",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.RemovePermissionFromRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToPermission{})
		cmdutil.HandleError(err, "failed to remove-permission-from-role")

		cmd.Println(cmdutil.FormatJSON(res, "permissions"))
	},
}

func init() {
	Cmd.AddCommand(RemovePermissionFromRoleCmd)
}
