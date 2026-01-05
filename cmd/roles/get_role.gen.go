package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRoleCmd = &cobra.Command{
	Use:     "get-role [role_id]",
	Aliases: []string{"get"},
	Short:   "Get a role",
	Long: `Get a role
Documentation: https://docs.datadoghq.com/api/latest/roles/#get-role`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RoleResponse
		var err error

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.GetRole(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-role")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {

	Cmd.AddCommand(GetRoleCmd)
}
