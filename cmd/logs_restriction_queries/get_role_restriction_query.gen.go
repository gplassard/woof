package logs_restriction_queries

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRoleRestrictionQueryCmd = &cobra.Command{
	Use:   "get-role-restriction-query [role_id]",
	Short: "Get restriction query for a given role",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.GetRoleRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-role-restriction-query: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_restriction_queries")
	},
}

func init() {
	Cmd.AddCommand(GetRoleRestrictionQueryCmd)
}
