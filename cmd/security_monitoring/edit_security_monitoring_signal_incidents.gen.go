package security_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var EditSecurityMonitoringSignalIncidentsCmd = &cobra.Command{
	Use:   "edit-security-monitoring-signal-incidents [signal_id]",
	Short: "Change the related incidents of a security signal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.EditSecurityMonitoringSignalIncidents(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityMonitoringSignalIncidentsUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to edit-security-monitoring-signal-incidents: %v", err)
		}

		cmdutil.PrintJSON(res, "signal_metadata")
	},
}

func init() {
	Cmd.AddCommand(EditSecurityMonitoringSignalIncidentsCmd)
}
