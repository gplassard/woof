package logs_restriction_queries

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRestrictionQueryCmd = &cobra.Command{
	Use:   "get-restriction-query [restriction_query_id]",
	
	Short: "Get a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.GetRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-restriction-query: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_restriction_queries")
	},
}

func init() {
	Cmd.AddCommand(GetRestrictionQueryCmd)
}
