package roles

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateRoleCmd = &cobra.Command{
	Use:   "create-role",
	Aliases: []string{ "create", },
	Short: "Create role",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.CreateRole(client.NewContext(apiKey, appKey, site), datadogV2.RoleCreateRequest{})
		cmdutil.HandleError(err, "failed to create-role")

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(CreateRoleCmd)
}
