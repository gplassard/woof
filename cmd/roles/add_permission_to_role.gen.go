package roles

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AddPermissionToRoleCmd = &cobra.Command{
	Use:     "add-permission-to-role [role_id]",
	Aliases: []string{"add-permission-to"},
	Short:   "Grant permission to a role",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.AddPermissionToRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToPermission{})
		cmdutil.HandleError(err, "failed to add-permission-to-role")

		cmd.Println(cmdutil.FormatJSON(res, "permissions"))
	},
}

func init() {
	Cmd.AddCommand(AddPermissionToRoleCmd)
}
