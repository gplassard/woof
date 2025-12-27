package roles

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddPermissionToRoleCmd = &cobra.Command{
	Use:   "addpermissiontorole [role_id]",
	Short: "Grant permission to a role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.AddPermissionToRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToPermission{})
		if err != nil {
			log.Fatalf("failed to addpermissiontorole: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(AddPermissionToRoleCmd)
}
