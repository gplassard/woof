package roles

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRolePermissionsCmd = &cobra.Command{
	Use:   "listrolepermissions [role_id]",
	Short: "List permissions for a role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.ListRolePermissions(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to listrolepermissions: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(ListRolePermissionsCmd)
}
