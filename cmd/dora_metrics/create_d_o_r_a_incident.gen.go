package dora_metrics

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDORAIncidentCmd = &cobra.Command{
	Use:   "create_d_o_r_a_incident",
	Short: "Send an incident event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateDORAIncident(client.NewContext(apiKey, appKey, site), datadogV2.DORAFailureRequest{})
		if err != nil {
			log.Fatalf("failed to create_d_o_r_a_incident: %v", err)
		}

		cmdutil.PrintJSON(res, "dora_failure")
	},
}

func init() {
	Cmd.AddCommand(CreateDORAIncidentCmd)
}
