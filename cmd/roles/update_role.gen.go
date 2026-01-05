package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateRoleCmd = &cobra.Command{
	Use:     "update-role [role_id]",
	Aliases: []string{"update"},
	Short:   "Update a role",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.UpdateRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RoleUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-role")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(UpdateRoleCmd)
}
