package logs_restriction_queries

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRestrictionQueryCmd = &cobra.Command{
	Use:   "deleterestrictionquery [restriction_query_id]",
	Short: "Delete a restriction query",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		_, err := api.DeleteRestrictionQuery(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deleterestrictionquery: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRestrictionQueryCmd)
}
