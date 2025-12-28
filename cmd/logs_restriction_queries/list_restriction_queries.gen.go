package logs_restriction_queries

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRestrictionQueriesCmd = &cobra.Command{
	Use:   "list-restriction-queries",
	
	Short: "List restriction queries",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.ListRestrictionQueries(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-restriction-queries")

		cmdutil.PrintJSON(res, "logs_restriction_queries")
	},
}

func init() {
	Cmd.AddCommand(ListRestrictionQueriesCmd)
}
