package apm_retention_filters

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteApmRetentionFilterCmd = &cobra.Command{
	Use:   "delete_apm_retention_filter [filter_id]",
	Short: "Delete a retention filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewAPMRetentionFiltersApi(client.NewAPIClient())
		_, err := api.DeleteApmRetentionFilter(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_apm_retention_filter: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteApmRetentionFilterCmd)
}
