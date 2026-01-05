package roles

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RemoveUserFromRoleCmd = &cobra.Command{
	Use:     "remove-user-from-role [role_id]",
	Aliases: []string{"remove-user-from"},
	Short:   "Remove a user from a role",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.RemoveUserFromRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToUser{})
		cmdutil.HandleError(err, "failed to remove-user-from-role")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {
	Cmd.AddCommand(RemoveUserFromRoleCmd)
}
