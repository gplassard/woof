package roles

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRoleUsersCmd = &cobra.Command{
	Use:   "list-role-users [role_id]",
	Aliases: []string{ "list-users", },
	Short: "Get all users of a role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.ListRoleUsers(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to list-role-users: %v", err)
		}

		cmdutil.PrintJSON(res, "users")
	},
}

func init() {
	Cmd.AddCommand(ListRoleUsersCmd)
}
