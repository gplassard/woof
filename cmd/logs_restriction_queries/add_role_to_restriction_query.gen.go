package logs_restriction_queries

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var AddRoleToRestrictionQueryCmd = &cobra.Command{
	Use:   "add-role-to-restriction-query [restriction_query_id]",
	
	Short: "Grant role to a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		_, err := api.AddRoleToRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RelationshipToRole{})
		cmdutil.HandleError(err, "failed to add-role-to-restriction-query")

		
	},
}

func init() {
	Cmd.AddCommand(AddRoleToRestrictionQueryCmd)
}
