package roles

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateRoleCmd = &cobra.Command{
	Use:   "createrole",
	Short: "Create role",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.CreateRole(client.NewContext(apiKey, appKey, site), datadogV2.RoleCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createrole: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(CreateRoleCmd)
}
