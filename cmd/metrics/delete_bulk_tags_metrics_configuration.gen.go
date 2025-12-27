package metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteBulkTagsMetricsConfigurationCmd = &cobra.Command{
	Use:   "delete_bulk_tags_metrics_configuration",
	Short: "Delete tags for multiple metrics",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.DeleteBulkTagsMetricsConfiguration(client.NewContext(apiKey, appKey, site), datadogV2.MetricBulkTagConfigDeleteRequest{})
		if err != nil {
			log.Fatalf("failed to delete_bulk_tags_metrics_configuration: %v", err)
		}

		cmdutil.PrintJSON(res, "metrics")
	},
}

func init() {
	Cmd.AddCommand(DeleteBulkTagsMetricsConfigurationCmd)
}
