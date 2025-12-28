package rum_retention_filters

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRetentionFilterCmd = &cobra.Command{
	Use:   "get-retention-filter [app_id] [rf_id]",
	
	Short: "Get a RUM retention filter",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		res, _, err := api.GetRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to get-retention-filter: %v", err)
		}

		cmdutil.PrintJSON(res, "retention_filters")
	},
}

func init() {
	Cmd.AddCommand(GetRetentionFilterCmd)
}
