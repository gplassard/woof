package logs_restriction_queries

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var RemoveRoleFromRestrictionQueryCmd = &cobra.Command{
	Use:   "remove-role-from-restriction-query [restriction_query_id]",
	
	Short: "Revoke role from a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		_, err := api.RemoveRoleFromRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToRole{})
		if err != nil {
			log.Fatalf("failed to remove-role-from-restriction-query: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(RemoveRoleFromRestrictionQueryCmd)
}
