package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetResourceEvaluationFiltersCmd = &cobra.Command{
	Use:   "get-resource-evaluation-filters",
	
	Short: "List resource filters",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetResourceEvaluationFilters(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-resource-evaluation-filters: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_resource_filter")
	},
}

func init() {
	Cmd.AddCommand(GetResourceEvaluationFiltersCmd)
}
