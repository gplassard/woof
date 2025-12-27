package rum_retention_filters

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRetentionFilterCmd = &cobra.Command{
	Use:   "deleteretentionfilter [app_id] [rf_id]",
	Short: "Delete a RUM retention filter",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumRetentionFiltersApi(client.NewAPIClient())
		_, err := api.DeleteRetentionFilter(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deleteretentionfilter: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRetentionFilterCmd)
}
