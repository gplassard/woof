package roles

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRolesCmd = &cobra.Command{
	Use:   "list-roles",
	Aliases: []string{ "list", },
	Short: "List roles",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.ListRoles(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-roles: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(ListRolesCmd)
}
