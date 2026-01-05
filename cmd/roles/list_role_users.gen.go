package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListRoleUsersCmd = &cobra.Command{
	Use:     "list-role-users [role_id]",
	Aliases: []string{"list-users"},
	Short:   "Get all users of a role",
	Long: `Get all users of a role
Documentation: https://docs.datadoghq.com/api/latest/roles/#list-role-users`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UsersResponse
		var err error

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.ListRoleUsers(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-role-users")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {

	Cmd.AddCommand(ListRoleUsersCmd)
}
