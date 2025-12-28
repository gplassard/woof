package dora_metrics

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDORAIncidentCmd = &cobra.Command{
	Use:   "create-d-o-r-ai-ncident",
	
	Short: "Send an incident event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateDORAIncident(client.NewContext(apiKey, appKey, site), datadogV2.DORAFailureRequest{})
		cmdutil.HandleError(err, "failed to create-d-o-r-ai-ncident")

		cmdutil.PrintJSON(res, "dora_failure")
	},
}

func init() {
	Cmd.AddCommand(CreateDORAIncidentCmd)
}
