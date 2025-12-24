package roles

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddUserToRoleCmd = &cobra.Command{
	Use:   "addusertorole [role_id]",
	Short: "Add a user to a role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.AddUserToRole(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToUser{})
		if err != nil {
			log.Fatalf("failed to addusertorole: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(AddUserToRoleCmd)
}
