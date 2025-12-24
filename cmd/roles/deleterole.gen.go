package roles

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRoleCmd = &cobra.Command{
	Use:   "deleterole [role_id]",
	Short: "Delete role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		_, err := api.DeleteRole(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deleterole: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRoleCmd)
}
