package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var RemoveUserFromRoleCmd = &cobra.Command{
	Use:     "remove-user-from-role [role_id] [payload]",
	Aliases: []string{"remove-user-from"},
	Short:   "Remove a user from a role",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.RelationshipToUser
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.RemoveUserFromRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to remove-user-from-role")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {
	Cmd.AddCommand(RemoveUserFromRoleCmd)
}
