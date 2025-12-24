package dora_metrics

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDORAIncidentCmd = &cobra.Command{
	Use:   "createdoraincident",
	Short: "Send an incident event",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDORAMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateDORAIncident(client.NewContext(apiKey, appKey, site), datadogV2.DORAFailureRequest{})
		if err != nil {
			log.Fatalf("failed to createdoraincident: %v", err)
		}

		cmdutil.PrintJSON(res, "dora_metrics")
	},
}

func init() {
	Cmd.AddCommand(CreateDORAIncidentCmd)
}
