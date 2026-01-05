package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var RemovePermissionFromRoleCmd = &cobra.Command{
	Use:     "remove-permission-from-role [role_id] [payload]",
	Aliases: []string{"remove-permission-from"},
	Short:   "Revoke permission",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.RelationshipToPermission
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.RemovePermissionFromRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to remove-permission-from-role")

		cmd.Println(cmdutil.FormatJSON(res, "permissions"))
	},
}

func init() {
	Cmd.AddCommand(RemovePermissionFromRoleCmd)
}
