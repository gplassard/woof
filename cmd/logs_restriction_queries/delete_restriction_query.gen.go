package logs_restriction_queries

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRestrictionQueryCmd = &cobra.Command{
	Use:   "delete-restriction-query [restriction_query_id]",
	
	Short: "Delete a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		_, err := api.DeleteRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-restriction-query")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRestrictionQueryCmd)
}
