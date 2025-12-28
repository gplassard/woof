package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListSecurityMonitoringHistsignalsCmd = &cobra.Command{
	Use:   "list_security_monitoring_histsignals",
	Short: "List hist signals",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListSecurityMonitoringHistsignals(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list_security_monitoring_histsignals: %v", err)
		}

		cmdutil.PrintJSON(res, "signal")
	},
}

func init() {
	Cmd.AddCommand(ListSecurityMonitoringHistsignalsCmd)
}
