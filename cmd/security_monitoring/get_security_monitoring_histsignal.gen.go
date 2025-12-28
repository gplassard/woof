package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetSecurityMonitoringHistsignalCmd = &cobra.Command{
	Use:   "get-security-monitoring-histsignal [histsignal_id]",
	Aliases: []string{ "get-histsignal", },
	Short: "Get a hist signal's details",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetSecurityMonitoringHistsignal(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-security-monitoring-histsignal: %v", err)
		}

		cmdutil.PrintJSON(res, "signal")
	},
}

func init() {
	Cmd.AddCommand(GetSecurityMonitoringHistsignalCmd)
}
