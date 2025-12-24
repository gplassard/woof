package roles

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/spf13/cobra"
)

var listRoles = &cobra.Command{
	Use:   "list",
	Short: "List Datadog roles",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		rolesApi := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err := rolesApi.ListRoles(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list roles: %v", err)
		}

		cmdutil.PrintJSON(res, "roles")
	},
}
