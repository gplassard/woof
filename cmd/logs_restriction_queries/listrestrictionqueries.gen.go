package logs_restriction_queries

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListRestrictionQueriesCmd = &cobra.Command{
	Use:   "listrestrictionqueries",
	Short: "List restriction queries",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.ListRestrictionQueries(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listrestrictionqueries: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_restriction_queries")
	},
}

func init() {
	Cmd.AddCommand(ListRestrictionQueriesCmd)
}
