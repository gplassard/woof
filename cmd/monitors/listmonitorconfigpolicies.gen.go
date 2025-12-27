package monitors

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListMonitorConfigPoliciesCmd = &cobra.Command{
	Use:   "listmonitorconfigpolicies",
	Short: "Get all monitor configuration policies",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.ListMonitorConfigPolicies(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listmonitorconfigpolicies: %v", err)
		}

		cmdutil.PrintJSON(res, "monitors")
	},
}

func init() {
	Cmd.AddCommand(ListMonitorConfigPoliciesCmd)
}
