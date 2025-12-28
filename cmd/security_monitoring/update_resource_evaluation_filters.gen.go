package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateResourceEvaluationFiltersCmd = &cobra.Command{
	Use:   "update_resource_evaluation_filters",
	Short: "Update resource filters",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateResourceEvaluationFilters(client.NewContext(apiKey, appKey, site), datadogV2.UpdateResourceEvaluationFiltersRequest{})
		if err != nil {
			log.Fatalf("failed to update_resource_evaluation_filters: %v", err)
		}

		cmdutil.PrintJSON(res, "csm_resource_filter")
	},
}

func init() {
	Cmd.AddCommand(UpdateResourceEvaluationFiltersCmd)
}
