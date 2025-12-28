package roles

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRoleCmd = &cobra.Command{
	Use:   "get-role [role_id]",
	Aliases: []string{ "get", },
	Short: "Get a role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.GetRole(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-role: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(GetRoleCmd)
}
