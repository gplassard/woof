package logs_restriction_queries

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ReplaceRestrictionQueryCmd = &cobra.Command{
	Use:   "replacerestrictionquery [restriction_query_id]",
	Short: "Replace a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.ReplaceRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RestrictionQueryUpdatePayload{})
		if err != nil {
			log.Fatalf("failed to replacerestrictionquery: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_restriction_queries")
	},
}

func init() {
	Cmd.AddCommand(ReplaceRestrictionQueryCmd)
}
