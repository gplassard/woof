package logs_restriction_queries

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateRestrictionQueryCmd = &cobra.Command{
	Use:   "createrestrictionquery",
	Short: "Create a restriction query",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewLogsRestrictionQueriesApi(client.NewAPIClient())
		res, _, err := api.CreateRestrictionQuery(client.NewContext(apiKey, appKey, site), datadogV2.RestrictionQueryCreatePayload{})
		if err != nil {
			log.Fatalf("failed to createrestrictionquery: %v", err)
		}

		cmdutil.PrintJSON(res, "logs_restriction_queries")
	},
}

func init() {
	Cmd.AddCommand(CreateRestrictionQueryCmd)
}
