package roles

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CloneRoleCmd = &cobra.Command{
	Use:     "clone-role [role_id]",
	Aliases: []string{"clone"},
	Short:   "Create a new role by cloning an existing role",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.CloneRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RoleCloneRequest{})
		cmdutil.HandleError(err, "failed to clone-role")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {
	Cmd.AddCommand(CloneRoleCmd)
}
