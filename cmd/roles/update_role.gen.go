package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateRoleCmd = &cobra.Command{
	Use:     "update-role [role_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update a role",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RoleUpdateResponse
		var err error

		var body datadogV2.RoleUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.UpdateRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-role")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(UpdateRoleCmd)
}
