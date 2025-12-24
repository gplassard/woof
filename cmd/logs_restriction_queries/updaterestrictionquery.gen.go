package logs_restriction_queries

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateRestrictionQueryCmd = &cobra.Command{
	Use:   "updaterestrictionquery [restriction_query_id]",
	Short: "Update a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.UpdateRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0], datadogV2.RestrictionQueryUpdatePayload{})
		if err != nil {
			log.Fatalf("failed to updaterestrictionquery: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_restriction_queries")
	},
}

func init() {
	Cmd.AddCommand(UpdateRestrictionQueryCmd)
}
