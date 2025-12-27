package roles

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RemovePermissionFromRoleCmd = &cobra.Command{
	Use:   "remove_permission_from_role [role_id]",
	Short: "Revoke permission",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.RemovePermissionFromRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToPermission{})
		if err != nil {
			log.Fatalf("failed to remove_permission_from_role: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(RemovePermissionFromRoleCmd)
}
