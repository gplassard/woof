package roles

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRoleTemplatesCmd = &cobra.Command{
	Use:   "listroletemplates",
	Short: "List role templates",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := api.ListRoleTemplates(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listroletemplates: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}

func init() {
	Cmd.AddCommand(ListRoleTemplatesCmd)
}
